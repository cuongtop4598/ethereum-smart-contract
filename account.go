package main

import (
	"context"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Interact with account
type UserAccount struct {
	Account common.Address
}

// GetBalanceCurrent is used to get the current balance of the account
func GetBalanceCurrent(client ethclient.Client, hexAccount string) *big.Int {
	account := common.HexToAddress(hexAccount)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	return balance
}

// GetBalanceAtBlock is used to get the current balance of the account at a given block
func GetBalanceAtBlock(client ethclient.Client, hexAccount string, blockNumber *big.Int) *big.Int {
	account := common.HexToAddress(hexAccount)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber) //
	if err != nil {
		log.Fatal(err)
	}
	return balanceAt
}

// GetFloatBalance is used to get the current balance of the ...
func GetFloatBalance(client ethclient.Client, hexAccount string, blockNumber *big.Int) *big.Float {
	fbalance := new(big.Float)
	account := common.HexToAddress(hexAccount)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue
}

func GetPendingBalance(client ethclient.Client, hexAccount string, blockNumber *big.Int) *big.Int {
	account := common.HexToAddress(hexAccount)
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}
	return pendingBalance
}
