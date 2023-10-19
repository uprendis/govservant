package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

func privateKeyFromMnemonic(mnemonic string, targetAddr common.Address) *ecdsa.PrivateKey {
	//println(mnemonic)
	// Generate a Bip32 HD wallet for the mnemonic and a user-supplied password.
	seed := bip39.NewSeed(mnemonic, "")
	//println(common.Bytes2Hex(seed))

	// Generate master key (m)
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		log.Fatalf("Error generating master key: %v", err)
	}
	//println(masterKey.String())

	// This is the maximum number of addresses you want to check. Be careful with large numbers as this might take a while.
	maxAddressesToCheck := 1000

	for i := 0; i < maxAddressesToCheck; i++ {
		// Derive the next account's key (BIP44: m/44'/60'/0'/0/i)
		path := accounts.DerivationPath{0x80000000 + 44, 0x80000000 + 60, 0x80000000 + 0, 0, uint32(i)}
		accountKey, err := derivePrivateKey(masterKey, path...)
		if err != nil {
			log.Fatalf("Error deriving private key at index %d: %v", i, err)
		}

		//println(accountKey.String())
		keyBytes, err := accountKey.ECPrivKey()
		if err != nil {
			log.Fatalf("Error getting EC private key: %v", err)
		}

		privateKey, err := crypto.ToECDSA(keyBytes.Serialize())
		if err != nil {
			log.Fatalf("Error creating ECDSA private key: %v", err)
		}

		address := crypto.PubkeyToAddress(privateKey.PublicKey)

		// Compare with the target address
		//println(address.String(), targetAddr.String())
		if address == targetAddr {
			fmt.Printf("Found matching address at index %d\n", i)
			fmt.Println("Address:", address.Hex())
			return privateKey
		}
	}

	return nil
}

// derivePrivateKey derives the private key of the derivation path.
func derivePrivateKey(masterKey *hdkeychain.ExtendedKey, path ...uint32) (*hdkeychain.ExtendedKey, error) {
	derivedKey := masterKey
	for _, n := range path {
		var err error
		derivedKey, err = derivedKey.Child(n)
		if err != nil {
			return nil, err
		}
	}
	return derivedKey, nil
}
