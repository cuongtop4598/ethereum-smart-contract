package main

import "github.com/ethereum/go-ethereum/common"

/**
* Accounts on ethereum are either wallet addresses or smart contract addresses.
* they look like 0x71c......
* they are what you use for sending ETH to another user and aslo are used for
* referring to a smart contract on the blockchain when needing to interact with it.
* they are unique and are derived from a private key.
**/

// in order to use account addresses with go-ethereum, you must first convert them to the go-ethereum common.Address type.

func NewAddress(hex string) common.Address {
	return common.HexToAddress(hex)
}
