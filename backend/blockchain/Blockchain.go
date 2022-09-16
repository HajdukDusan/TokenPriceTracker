package blockchain

import (
	"backendtask/api"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client
var PriceSetterContract Contract
var PrivateKey string

func Connect(rpcUrl string, contractAddress string, privateKey string) {

	var err error

	// connect to rpc node
	Client, err = ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}

	address := common.HexToAddress(contractAddress)

	// creating api object to intract with smart contract function
	api, err := api.NewApi(address, Client)
	if err != nil {
		panic(err)
	}

	PriceSetterContract = Contract{api}

	PrivateKey = privateKey
}

