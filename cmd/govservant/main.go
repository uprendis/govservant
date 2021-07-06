package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func isHelp(args []string) bool {
	return len(os.Args) >= 2 && (os.Args[1] == "-h" || os.Args[1] == "-help" || os.Args[1] == "--help" || os.Args[1] == "help")
}

func main() {
	if len(os.Args) != (1+4) || isHelp(os.Args) {
		fmt.Printf("Usage: %s /path/to/privkey /path/to/password governance_address RPC_URL\n", os.Args[0])
		return
	}
	// parse arguments
	privkeyPath := os.Args[1]
	passwordPath := os.Args[2]
	if !common.IsHexAddress(os.Args[3]) {
		fmt.Fprintf(os.Stderr, "Governance address parsing error\n")
		return
	}
	govAddress := common.HexToAddress(os.Args[3])
	rpcUrl := os.Args[4]

	// decrypt the private key

	keyjson, err := ioutil.ReadFile(privkeyPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Private key reading error: %v\n", err)
		return
	}
	// read first line of the password file
	authText, err := ioutil.ReadFile(passwordPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Password file reading error: %v\n", err)
		return
	}
	auth := strings.Split(string(authText), "\n")[0]

	key, err := keystore.DecryptKey(keyjson, auth)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Private key decrypting error: %v\n", err)
		return
	}

	signerFn := func(s types.Signer, _ common.Address, tx *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(tx, s, key.PrivateKey)
	}

	// Governance maintenance

	// Watch for Ctrl-C while the import is running.
	// If a signal is received, the import will stop.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(interrupt)
	for {
		handled, err := maintainGovTasks(key.Address, govAddress, signerFn, rpcUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Governance maintainance error: %v\n", err)
		}
		fmt.Printf("Handled %d tasks\n", handled)
		timer := time.NewTimer(30 * time.Second)
		select {
		case <-timer.C:
			continue
		case <-interrupt:
			return
		}
	}
}
