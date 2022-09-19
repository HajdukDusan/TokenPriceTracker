package main

import (
	"math/big"
	"os"

	"backendtask/blockchain"
	"backendtask/config"
	"backendtask/routines"
	"backendtask/server"

	"github.com/joho/godotenv"
)



func main() {

	// Load Env File
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	contractAddress := os.Getenv("PRICE_SETTER_ADDRESS")
	rpcUrl := os.Getenv("GOERLI_RPC_WS_URL")
	privateKey := os.Getenv("PRIVATE_KEY")

	// try to initialize the client
	blockchain.Connect(rpcUrl)
	if err != nil {
		panic(err)
	}

	// initialize the priceSetter api
	priceSetterContract, err := blockchain.CreateContractAPI(contractAddress)
	if err != nil {
		panic(err)
	}

	// create a map of coins from all symbols in config
	var coins = make(map[string]*big.Int)
	for _, symbol := range config.Symbols {
		coins[symbol] = nil
	}

	// launch update goroutine
	// maps are passed by ref so no need for pointers
	go routines.UpdatePrices(coins, priceSetterContract, privateKey)

	server.Start(priceSetterContract)
}
