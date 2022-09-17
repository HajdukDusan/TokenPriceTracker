package blockchain

import (
	"backendtask/api"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client
var PriceSetterContract Contract
var PrivateKey string

// Connect to rpc node and initialize client
func Connect(rpcUrl string, contractAddress string, privateKey string) error {

	var err error

	Client, err = ethclient.Dial(rpcUrl)
	if err != nil {
		return err
	}

	address := common.HexToAddress(contractAddress)

	// creating api object to interact with smart contract functions
	api, err := api.NewApi(address, Client)
	if err != nil {
		return err
	}

	// wrap the api into Contract struct
	PriceSetterContract = Contract{api}

	PrivateKey = privateKey

	return nil
}

