package blockchain

import (
	"backendtask/api"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type Contract struct {
	*api.Api
}

func (contract *Contract) GetSymbolPrice(symbol string) (*big.Int, error) {
	return contract.PriceOf(&bind.CallOpts{}, symbol)
}