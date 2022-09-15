package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"time"

	"backendtask/api" // generated smart contract bindings
	"backendtask/blockchain"
	"backendtask/config"
	"backendtask/server"
	"backendtask/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var coins = map[string]int64{
	"ethereum": 0,
	"bitcoin":  0,
}

func updatePrices(client *ethclient.Client, contract *api.Api, privateKey string) {
	//https://api.coingecko.com/api/v3/coins/bitcoin/history?date=01-01-2020&localization=false
	for true {

		ids := ""
		index := 1
		for key := range coins {
			ids += key

			if index != len(coins) {
				ids += ","
			}
			index++
		}

		response, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=" + ids + "&vs_currencies=usd")
		if err != nil {
			panic(err)
		}

		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		response.Body.Close()

		//Convert the body to type string
		res := string(body)

		resBytes := []byte(res)                  // Converting the string "res" into byte array
		var jsonRes map[string]interface{}       // declaring a map for key names as string and values as interface{}
		err = json.Unmarshal(resBytes, &jsonRes) // Unmarshalling

		if err != nil {
			panic(err)
		}

		// Type casting again so that interface{} -> map[string]interface{}

		for key := range coins {
			coin := jsonRes[key].(map[string]interface{}) // type the interface again to a map with key string type and value as interface
			price := coin["usd"].(float64)

			result, err := utils.ScaleFloatToInt(price)
			if err != nil {
				fmt.Println(err)
				continue
			}

			coins[key] = result
		}

		for key, value := range coins {

			auth := getAccountAuth(client, privateKey)

			tx, err := contract.Set(auth, key, big.NewInt(value))
			if err != nil {
				fmt.Println("Failed to send tx!")
				fmt.Println(err)
			} else {
				fmt.Printf("Tx successfully sent: %s\n", tx.Hash().Hex())
			}

		}

		// stop thread for config.UpdateInterval seconds
		time.Sleep(time.Duration(config.UpdateInterval) * time.Second)
	}
}

func main() {

	// Load Env File
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	contractAddress := os.Getenv("PRICE_SETTER_ADDRESS")


	// address := common.HexToAddress("0xCf72bbDc50e2a360F175441D9748271DFd4DC3AA")

	//creating api object to intract with smart contract function
	// contract, err := api.NewApi(address, client)
	// if err != nil {
	// 	panic(err)
	// }

	// currBlock, err := client.BlockNumber(context.Background())
	// if err != nil {
	// 	panic(err)
	// }

	// fetchEvents(client, 0, int64(currBlock))

	// privateKey := os.Getenv("PRIVATE_KEY")
	// launch goroutines
	// go updatePrices(client, contract, privateKey)

	blockchain.Connect("https://eth-goerli.g.alchemy.com/v2/dMsipp8n2rbpCS9-6ixSDYijn1zvLsBC")

	server.Start(contractAddress)
}

// function to create auth for any account from its private key
func getAccountAuth(client *ethclient.Client, privateKeyAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(privateKeyAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Suggested Gas Price: " + gasPrice.String())

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units

	// gasPrice x2 for fast mining
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(2))

	return auth
}
