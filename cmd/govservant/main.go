package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
)

func isHelp(args []string) bool {
	return len(os.Args) >= 2 && (os.Args[1] == "-h" || os.Args[1] == "-help" || os.Args[1] == "--help" || os.Args[1] == "help")
}

func main() {
	if len(os.Args) != (1+4) || isHelp(os.Args) {
		fmt.Printf("Usage: %s /path/to/funder-mnemonic /path/to/stolen-mnemonic destination-addr RPC_URL\n", os.Args[0])
		return
	}
	// parse arguments
	funderMnemonicPath := os.Args[1]
	stolenMnemonicPath := os.Args[2]
	if !common.IsHexAddress(os.Args[3]) {
		fmt.Fprintf(os.Stderr, "Destination address parsing error\n")
		return
	}
	destinationAddr := common.HexToAddress(os.Args[3])
	rpcUrl := os.Args[4]

	// decrypt the private key

	funderMnemonic, err := ioutil.ReadFile(funderMnemonicPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Funder mnemonic reading error: %v\n", err)
		return
	}
	stolenMnemonic, err := ioutil.ReadFile(stolenMnemonicPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Stolen mnemonic reading error: %v\n", err)
		return
	}
	funderSplit := strings.Split(string(funderMnemonic), "->")
	if len(funderSplit) != 2 {
		fmt.Fprintf(os.Stderr, "Mnemonic must have format: addr -> mnemonic-words\n", err)
		return
	}
	stolenSplit := strings.Split(string(stolenMnemonic), "->")
	if len(stolenSplit) != 2 {
		fmt.Fprintf(os.Stderr, "Mnemonic must have format: addr -> mnemonic-words\n", err)
		return
	}
	if !common.IsHexAddress(strings.TrimSpace(funderSplit[0])) {
		fmt.Fprintf(os.Stderr, "Funder address parsing error\n")
		return
	}
	funderAddr := common.HexToAddress(strings.TrimSpace(funderSplit[0]))
	if !common.IsHexAddress(strings.TrimSpace(stolenSplit[0])) {
		fmt.Fprintf(os.Stderr, "Stolen address parsing error\n")
		return
	}
	stolenAddr := common.HexToAddress(strings.TrimSpace(stolenSplit[0]))

	funderKey := privateKeyFromMnemonic(strings.TrimSpace(funderSplit[1]), funderAddr)
	if funderKey == nil {
		fmt.Fprintf(os.Stderr, "Failed to derive funder key\n")
		return
	}
	stolenKey := privateKeyFromMnemonic(strings.TrimSpace(stolenSplit[1]), stolenAddr)
	if funderKey == nil {
		fmt.Fprintf(os.Stderr, "Failed to derive stolen key\n")
		return
	}

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
