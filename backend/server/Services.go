package server

import (
	"backendtask/blockchain"
	"backendtask/dto"
	"backendtask/utils"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func FetchDTOEvents(contractAddress string, fromTimestamp *big.Int, toTimestamp *big.Int, symbols []string) ([]dto.PriceChangeEventDTO, error) {

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

	eventsDTO := make([]dto.PriceChangeEventDTO, 0)


	for _, event := range events {

		// skip events that are not in timestamp range
		if event.Timestamp.Cmp(fromTimestamp) == -1 {
			continue
		}
		if event.Timestamp.Cmp(toTimestamp) == 1 {
			continue
		}

		priceFloat, err := utils.ScaleIntToFloat(event.Price.Int64())
		if err != nil {
			return nil, err
		}

		eventsDTO = append(
			eventsDTO, dto.PriceChangeEventDTO {
			BlockNumber: event.BlockNumber,
			Symbol:      symbolHashMap[event.SymbolHash],	// get symbol from symbolHashMap with symbol hash as key
			Price:       priceFloat,
			Timestamp:   event.Timestamp.Uint64(),
		})
	}

	return eventsDTO, nil
}