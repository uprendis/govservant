package main

import (
	"math/big"
	"testing"
)

func Test1(t *testing.T) {
	balance := big.NewInt(510000000000000000)
	if balance.Sign() == 0 {
		return
	}
	println("non-zero token balance", balance.String())
	price := 0.0005
	decimals := uint8(18)

	value := new(big.Int).Mul(balance, big.NewInt(int64(price*1000000)))
	for i := uint8(0); i < decimals; i++ {
		value.Div(value, big.NewInt(10))
	}
	println(value.String())
}
