package dto

type PriceChangeEventDTO struct {
	BlockNumber uint64
	Timestamp   uint64
	Symbol      string
	Price       float64
}