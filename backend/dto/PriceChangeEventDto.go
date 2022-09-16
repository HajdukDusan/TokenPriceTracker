package dto

type PriceChangeEventDTO struct {
	BlockNumber uint64
	Symbol      string
	Price       float64
}