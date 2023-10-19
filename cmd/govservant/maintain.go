package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/uprendis/govservant/erc20"
)

func fetchPrice(network, token string) (float64, error) {
	//return 0.5, nil
	// This is a dummy URL; you'd replace this with the actual endpoint from a service like CoinGecko.
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/token_price/%s?contract_addresses=%s&vs_currencies=usd&precision=6", network, token)

	// Make a request to the external API.
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// Here, you would parse the JSON response based on the structure of the external API's response.
	// This structure depends on the API's response format.
	var result map[string]map[string]float64
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, errors.Wrap(err, string(body))
	}

	return result[strings.ToLower(token)]["usd"], nil // Be sure to replace "usd" with the actual currency you're comparing against.
}

var ongoingErc20Transfer int64

func handleErc20(sender, funder, recipient common.Address, senderKey, funderKey *ecdsa.PrivateKey, rpcUrl string, quit chan struct{}) (handled *big.Int, err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Minute)
	defer cancel()
	client, err := ethclient.DialContext(ctx, rpcUrl)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	network := ""
	if chainID.Uint64() == 1 {
		network = "ethereum"
	} else if chainID.Uint64() == 250 {
		network = "fantom"
	} else {
		return nil, errors.New("not eth/ftm mainnet network")
	}

	signerFn := func(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(tx, types.LatestSignerForChainID(chainID), senderKey)
	}

	blockNum, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	topics := [][]common.Hash{{}, {}, {sender.Hash()}}

	ch := make(chan types.Log, 10000)
	sub, err := client.SubscribeFilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(int64(blockNum) - 100),
		ToBlock:   nil,
		Addresses: nil,
		Topics:    topics,
	}, ch)
	if err != nil {
		return nil, err
	}
	timer := time.NewTimer(4 * time.Minute + time.Second)

	callOpts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	for {
		select {
		case msg := <-ch:
			tokenAddr := msg.Address

			erc20, err := erc20.NewErc20(tokenAddr, client)
			if err != nil {
				return handled, err
			}
			balance, err := erc20.BalanceOf(callOpts, sender)
			if err != nil {
				return handled, err
			}
			if balance.Sign() == 0 {
				continue
			}
			println("detected non-zero token balance", tokenAddr.String(), balance.String())
			gasPrice, err := client.SuggestGasPrice(ctx)
			if err != nil {
				return handled, err
			}
			gasPrice.Mul(gasPrice, common.Big2)
			price, err := fetchPrice(network, tokenAddr.String())
			if err != nil {
				return handled, err
			}
			decimals, err := erc20.Decimals(callOpts)
			if err != nil {
				return handled, err
			}
			value := new(big.Int).Mul(balance, big.NewInt(int64(price*1000000)))
			for i := uint8(0); i < decimals; i++ {
				value.Div(value, big.NewInt(10))
			}
			//println(value.String())
			if value.Cmp(big.NewInt(2 * 1000000)) < 0 {
				continue
			}
			println("send non-zero token balance", tokenAddr.String(), balance.String())

			nonce, err := client.NonceAt(ctx, sender, nil)
			if err != nil {
				return handled, err
			}
			nonceFunder, err := client.NonceAt(ctx, funder, nil)
			if err != nil {
				return handled, err
			}
			transactOpts := &bind.TransactOpts{
				From:     sender,
				Nonce:    new(big.Int).SetUint64(nonce),
				Signer:   signerFn,
				GasPrice: gasPrice,
				GasLimit: 120000,
				Context:  ctx,
				NoSend:   true,
			}
			tx2, err := erc20.Transfer(transactOpts, recipient, balance)
			if err != nil {
				return handled, err
			}
			tx1 := types.NewTx(&types.LegacyTx{
				Nonce:    nonceFunder,
				To:       &sender,
				Value:    new(big.Int).Mul(big.NewInt(120000), gasPrice),
				Gas:      21000,
				GasPrice: gasPrice,
			})
			tx1, err = types.SignTx(tx1, types.LatestSignerForChainID(chainID), funderKey)
			if err != nil {
				return handled, err
			}
			balance2, err := erc20.BalanceOf(callOpts, sender)
			if err != nil {
				return handled, err
			}
			if balance2.Cmp(balance) < 0 {
				continue
			}
			atomic.StoreInt64(&ongoingErc20Transfer, time.Now().UnixNano())
			go func() {
				for i := 0; i < 30; i++ {
					client.SendTransaction(ctx, tx1)
					time.Sleep(time.Millisecond * 10)
					client.SendTransaction(ctx, tx2)
					time.Sleep(time.Millisecond * 500)
				}
			}()
			if handled == nil {
				handled = value
			} else {
				handled.Add(handled, value)
			}
		case <-quit:
			sub.Unsubscribe()
			return handled, nil
		case <-timer.C:
			sub.Unsubscribe()
			return handled, nil
		}
	}
}

func handleFtm(sender, recipient common.Address, senderKey *ecdsa.PrivateKey, rpcUrl string, quit chan struct{}) (handled *big.Int, err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Minute)
	defer cancel()
	client, err := ethclient.DialContext(ctx, rpcUrl)
	if err != nil {
		return handled, err
	}
	defer client.Close()
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return handled, err
	}

	ch := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(ctx, ch)
	if err != nil {
		return handled, err
	}
	timer := time.NewTimer(4 * time.Minute)

	for {
		select {
		case <-ch:
			prevErc20 := atomic.LoadInt64(&ongoingErc20Transfer)
			if time.Since(time.Unix(prevErc20/int64(time.Second), prevErc20%int64(time.Second))) < 10*time.Second {
				continue
			}
			balance, err := client.BalanceAt(ctx, sender, nil)
			if err != nil {
				return handled, err
			}

			gasPrice, err := client.SuggestGasPrice(ctx)
			if err != nil {
				return handled, err
			}
			gasPrice.Mul(gasPrice, common.Big2)

			fee := new(big.Int).Mul(big.NewInt(21000), gasPrice)
			if balance.Cmp(fee) < 0 {
				continue
			}
			println("send non-zero balance", balance.String())

			nonce, err := client.NonceAt(ctx, sender, nil)
			if err != nil {
				return handled, err
			}
			tx := types.NewTx(&types.LegacyTx{
				Nonce:    nonce,
				To:       &recipient,
				Value:    new(big.Int).Sub(balance, fee),
				Gas:      21000,
				GasPrice: gasPrice,
			})
			tx, err = types.SignTx(tx, types.LatestSignerForChainID(chainID), senderKey)
			if err != nil {
				return handled, err
			}
			go func() {
				for i := 0; i < 30; i++ {
					client.SendTransaction(ctx, tx)
					time.Sleep(time.Millisecond * 500)
				}
			}()
			if handled == nil {
				handled = new(big.Int).Sub(balance, fee)
			} else {
				handled.Add(handled, new(big.Int).Sub(balance, fee))
			}
		case <-quit:
			sub.Unsubscribe()
			return handled, nil
		case <-timer.C:
			sub.Unsubscribe()
			return handled, nil
		}
	}
}
