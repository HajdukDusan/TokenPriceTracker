package blockchain

import (
	"backendtask/api"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type Contract struct {
	*api.Api
}

func (contract *Contract) GetSymbolPrice(symbol string) (*big.Int, error) {
	return contract.PriceOf(&bind.CallOpts{}, symbol)
}

func (contract *Contract) SetSymbolPrice(auth *bind.TransactOpts, symbol string, price *big.Int) (*types.Transaction, error) {
	return contract.Set(auth, symbol, price)
}