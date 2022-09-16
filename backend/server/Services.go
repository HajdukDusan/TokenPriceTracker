package server

import (
	"backendtask/blockchain"
	"backendtask/dto"
	"backendtask/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func FetchDTOEvents(contractAddress string, fromTimestamp int64, toTimestamp int64, symbols []string) ([]dto.PriceChangeEventDTO, error) {

	// TODO TIMESTAMPS

	indexedValues := make([]common.Hash, len(symbols))
	symbolHashMap := make(map[string]string)

	// hash the array of symbols because they are indexed in events
	for indx := range symbols {
		hash := crypto.Keccak256Hash([]byte(symbols[indx]))
		indexedValues[indx] = hash
		symbolHashMap[hash.Hex()] = symbols[indx]
	}

	events, err := blockchain.FetchEvents(
		[]common.Address{common.HexToAddress(contractAddress)},
		nil,
		nil,
		[][]common.Hash{
			{},
			indexedValues,
		},
	)
	if err != nil {
		return nil, err
	}

	eventsDTO := make([]dto.PriceChangeEventDTO, len(events))

	// get symbol from symbolHashMap with symbol hash as key
	for indx, event := range events {

		priceFloat, err := utils.ScaleIntToFloat(event.Price.Int64())
		if err != nil {
			return nil, err
		}

		eventsDTO[indx] = dto.PriceChangeEventDTO{
			BlockNumber: event.BlockNumber,
			Symbol:      symbolHashMap[event.SymbolHash],
			Price:       priceFloat,
		}
	}

	return eventsDTO, nil
}