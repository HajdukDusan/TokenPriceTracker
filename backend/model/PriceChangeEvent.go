package model

import (
	"math/big"
)

type PriceChangeEvent struct {
	BlockNumber uint64
	SymbolHash  string
	Price       *big.Int
	Timestamp	*big.Int
}