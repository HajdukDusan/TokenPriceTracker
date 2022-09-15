package blockchain

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client

func Connect(rpcUrl string) {

	var err error

	// connect to rpc node
	Client, err = ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}

	// address := common.HexToAddress("0xCf72bbDc50e2a360F175441D9748271DFd4DC3AA")

	// creating api object to intract with smart contract function
	// contract, err := api.NewApi(address, client)
	// if err != nil {
	// 	panic(err)
	// }
}
