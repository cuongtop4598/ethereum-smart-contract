package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/cuongtop4598/Go-Ethereum/client/bindings"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAccountFromKeyStore() accounts.Account {
	file := "./tmp/UTC--2022-02-09T23-55-36.232904815Z--1f7b68c8c5d5cf6e47d47240afcc3a7d86b83046"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	password := "123456"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		fmt.Println(err)
	}

	return account
}

func CreateKey() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "123456"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
}

func main() {
	client, err := ethclient.Dial("http://27.72.105.169:10198")
	if err != nil {
		log.Fatal(err)
	}
	// get one account
	account := GetAccountFromKeyStore()
	fmt.Println("account address:", account.Address)
	// create binding options
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)

	// unlock account
	password := "123456"
	err = ks.Unlock(account, password)
	if err != nil {
		log.Fatal(err)
	}

	// deploy Tokenbalance contract =================================================
	{
		// tokenAddresses := []string{}
		// i := 0
		// for i < 3000 {

		// 	chainID := int64(451998)
		// 	auth, err := bind.NewKeyStoreTransactorWithChainID(ks, account, big.NewInt(chainID))
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fromAddress := account.Address
		// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	// gasPrice, err := client.SuggestGasPrice(context.Background())
		// 	// if err != nil {
		// 	// 	log.Fatal(err)
		// 	// }

		// 	auth.From = account.Address
		// 	auth.Nonce = big.NewInt(int64(nonce))
		// 	auth.Value = big.NewInt(0)       // in wei
		// 	auth.GasLimit = uint64(0x47b760) // in units gas limit: 134217728
		// 	auth.GasPrice = big.NewInt(10000)

		// 	tokenaddress, _, _, err := bindings.DeployERC20Basic(auth, client, big.NewInt(1000000000000000000)) // token
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println("token address", tokenaddress)
		// 	i++
		// 	tokenAddresses = append(tokenAddresses, tokenaddress.Hex())
		// }
		// data, _ := json.MarshalIndent(tokenAddresses, "", " ")
		// _ = ioutil.WriteFile("token_address.json", data, 6044)
	}

	// deploy document contract
	{
		// chainID := int64(451998)
		// auth, err := bind.NewKeyStoreTransactorWithChainID(ks, account, big.NewInt(chainID))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fromAddress := account.Address
		// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// // gasPrice, err := client.SuggestGasPrice(context.Background())
		// // if err != nil {
		// // 	log.Fatal(err)
		// // }

		// auth.From = account.Address
		// auth.Nonce = big.NewInt(int64(nonce))
		// auth.Value = big.NewInt(0)       // in wei
		// auth.GasLimit = uint64(0x47b760) // in units gas limit: 134217728
		// auth.GasPrice = big.NewInt(10000)

		// documentAddress, _, _, err := bindings.DeployDocument(auth, client)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// fmt.Println("contract address", documentAddress)
		// fmt.Println("Deploy done!")
	}

	//create contract instance ============================================================

	{

		contractAddress := common.HexToAddress("0x697FCA08C6F50c8aEc52eA434857d73F2308926A")
		contract, err := bindings.NewDePocketBridge(contractAddress, client)
		if err != nil {
			log.Fatal(err)
		}
		userAddress := common.HexToAddress("0x12A065612DC8cD96c1Ad67292f35C5b33F2f9807")

		addressStr := []string{}

		f, err := os.Open("./token_address.json")
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}
		_ = json.Unmarshal(data, &addressStr)

		commonAddresses := []common.Address{}
		for _, add := range addressStr {
			commonAddresses = append(commonAddresses, common.HexToAddress(add))
		}

		amount, err := contract.AllBalancesForManyAccounts(&bind.CallOpts{}, []common.Address{userAddress}, commonAddresses)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(amount)
	}

}
