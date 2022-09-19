package dto

import "math/big"

type PriceChangeEventDTO struct {
	BlockNumber uint64   `json:"BlockNumber"`
	Timestamp   *big.Int `json:"Timestamp"`
	Symbol      string   `json:"Symbol"`
	Price       float64  `json:"Price"`
}