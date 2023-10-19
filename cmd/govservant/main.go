package main

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

func isHelp(args []string) bool {
	return len(os.Args) >= 2 && (os.Args[1] == "-h" || os.Args[1] == "-help" || os.Args[1] == "--help" || os.Args[1] == "help")
}

func readKey(keyPath, passPath string) (common.Address, *ecdsa.PrivateKey) {
	keyjson, err := ioutil.ReadFile(keyPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Private key reading error: %v\n", err)
		os.Exit(1)
		return common.Address{}, nil
	}
	// read first line of the password file
	authText, err := ioutil.ReadFile(passPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Password file reading error: %v\n", err)
		os.Exit(1)
		return common.Address{}, nil
	}
	auth := strings.Split(string(authText), "\n")[0]

	key, err := keystore.DecryptKey(keyjson, auth)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Private key decrypting error: %v\n", err)
		os.Exit(1)
		return common.Address{}, nil
	}
	return key.Address, key.PrivateKey
}

func main() {
	if len(os.Args) != (1+6) || isHelp(os.Args) {
		fmt.Printf("Usage: %s /path/to/funder-key /path/to/funder-pass /path/to/stolen-key /path/to/stolen-pass destination-addr RPC_URL\n", os.Args[0])
		return
	}
	// parse arguments
	funderKeyPath := os.Args[1]
	funderPassPath := os.Args[2]
	stolenKeyPath := os.Args[3]
	stolenPassPath := os.Args[4]
	if !common.IsHexAddress(os.Args[5]) {
		fmt.Fprintf(os.Stderr, "Destination address parsing error\n")
		return
	}
	destinationAddr := common.HexToAddress(os.Args[3])
	rpcUrl := os.Args[6]

	// decrypt the private key

	funderAddr, funderKey := readKey(funderKeyPath, funderPassPath)
	stolenAddr, stolenKey := readKey(stolenKeyPath, stolenPassPath)

	// Governance maintenance

	// Watch for Ctrl-C while the import is running.
	// If a signal is received, the import will stop.
	interrupt := make(chan os.Signal, 1)
	quit := make(chan struct{})
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(interrupt)
	go func() {
		for {
			handled, err := handleErc20(stolenAddr, funderAddr, destinationAddr, stolenKey, funderKey, rpcUrl, quit)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Governance maintainance error: %v\n", err)
			}
			fmt.Printf("Handled ERC20 %s / 10^6 USD \n", handled.String())
			select {
			case <-quit:
				return
			default:
			}
		}
	}()
	go func() {
		for {
			handled, err := handleFtm(stolenAddr, destinationAddr, stolenKey, rpcUrl, quit)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Governance maintainance error: %v\n", err)
			}
			fmt.Printf("Handled %s wei \n", handled.String())
			select {
			case <-quit:
				return
			default:
			}
		}
	}()
	select {
	case <-interrupt:
		close(quit)
		return
	}
}
