package main

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/uprendis/govservant/govabi"
)

func maintainGovTasks(sender, govAddr common.Address, signer bind.SignerFn, rpcUrl string) (handled uint64, err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5 * time.Minute)
	defer cancel()
	client, err := ethclient.DialContext(ctx, rpcUrl)
	if err != nil {
		return handled, err
	}
	defer client.Close()
	nonce, err := client.NonceAt(ctx, sender, nil)
	if err != nil {
		return handled, err
	}

	gov, err := govabi.NewContract(govAddr, client)
	if err != nil {
		return handled, err
	}
	callOpts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	transactOpts := &bind.TransactOpts{
		From:    sender,
		Signer:  signer,
		Nonce:   new(big.Int).SetUint64(nonce),
		Context: ctx,
	}
	// handle tasks
	tasksCount, err := gov.TasksCount(callOpts)
	if err != nil {
		return handled, err
	}
	tasksCountU64 := tasksCount.Uint64()
	for i := uint64(0); i < tasksCountU64; i++ {
		_, err := gov.HandleTasks(transactOpts, new(big.Int).SetUint64(i), common.Big1)
		if err != nil {
			continue
		}
		nonce++
		transactOpts.Nonce = new(big.Int).SetUint64(nonce)
		handled++
	}
	// erase tasks
	_, _ = gov.TasksCleanup(transactOpts, big.NewInt(16))

	return handled, nil
}
