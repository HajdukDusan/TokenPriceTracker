package blockchain

import (
	"backendtask/api"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client

// Connect to rpc node and initialize client
func Connect(rpcUrl string) error {

	var err error

	Client, err = ethclient.Dial(rpcUrl)
	if err != nil {
		return err
	}

	return nil
}

// returns a wrapper class to interact with smart contract functions
func CreateContractAPI(contractAddress string) (*Contract, error) {
	address := common.HexToAddress(contractAddress)

	api, err := api.NewApi(address, Client)
	if err != nil {
		return nil, err
	}

	// wrap the api into Contract struct
	return &Contract{contractAddress, api}, nil
}

