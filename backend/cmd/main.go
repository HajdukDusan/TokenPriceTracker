package main

import (
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
	rpcUrl := os.Getenv("GOERLI_RPC_URL")
	privateKey := os.Getenv("PRIVATE_KEY")

	// try to initialize the client
	blockchain.Connect(rpcUrl, contractAddress, privateKey)
	if err != nil {
		panic(err)
	}

	// create a map of coins from all symbols in config
	var coins = make(map[string]int64)
	for _, symbol := range config.Symbols {
		coins[symbol] = 0
	}

	// launch update goroutine
	// maps are passed by ref so no need for pointers
	go routines.UpdatePrices(coins)

	server.Start(contractAddress)
}
