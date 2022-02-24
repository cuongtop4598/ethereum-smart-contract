package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/cuongtop4598/Go-Ethereum/client/bindings"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Deploy(fromAddress common.Address, client *ethclient.Client, account accounts.Account, keystore *keystore.KeyStore, chainID int64) {
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyStoreTransactorWithChainID(keystore, account, big.NewInt(chainID))
	if err != nil {
		log.Fatal(err)
	}
	auth.From = account.Address
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(8000000) // in units gas limit: 134217728
	auth.GasPrice = gasPrice
	address, tx, instance, err := bindings.DeployDePocketBridge(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println("Deploy Contract Done!")
	_ = instance
}
