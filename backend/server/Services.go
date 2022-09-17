package server

import (
	"backendtask/blockchain"
	"backendtask/dto"
	"backendtask/utils"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func FetchSymbolDTOPrice(symbol string, priceSetter *blockchain.Contract) (*dto.PriceDTO, error) {

	// get symbol price from contract
	contractPrice, err := priceSetter.GetSymbolPrice(symbol)
	if err != nil {
		return nil, err
	}
	floatContractPrice, err := utils.ScaleIntToFloat(contractPrice.Int64())
	if err != nil {
		return nil, err
	}

	coin := map[string]int64{}
	coin[symbol] = 0

	// get symbol price from API
	jsonResult, err := GetCoinGeckoPricesForCoins(coin)
	if err != nil {
		return nil, err
	}

	// get price from json
	result := jsonResult[symbol].(map[string]interface{})
	floatApiPrice := result["usd"].(float64)

	return &dto.PriceDTO{
		ApiPrice: floatApiPrice,
		ContractPrice: floatContractPrice,
	}, nil
}

// Fetch DTO events for symbols between timestamps
func FetchDTOEvents(contractAddress string, fromTimestamp *big.Int, toTimestamp *big.Int, symbols []string) ([]dto.PriceChangeEventDTO, error) {

	// for loading the symbol hash values
	indexedValues := make([]common.Hash, len(symbols))

	// for retrieving the symbol from its hash value
	symbolHashMap := make(map[string]string)

	for indx := range symbols {
		hash := crypto.Keccak256Hash([]byte(symbols[indx]))

		// hash the array of symbols because they are indexed in events
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

	// parse event to DTO
	eventsDTO := make([]dto.PriceChangeEventDTO, 0)
	for _, event := range events {

		// filter events that are not in timestamp range
		if event.Timestamp.Cmp(fromTimestamp) == -1 {
			continue
		}
		if event.Timestamp.Cmp(toTimestamp) == 1 {
			continue
		}

		// scale to float representation
		priceFloat, err := utils.ScaleIntToFloat(event.Price.Int64())
		if err != nil {
			return nil, err
		}

		eventsDTO = append(
			eventsDTO, dto.PriceChangeEventDTO{
				BlockNumber: event.BlockNumber,
				Symbol:      symbolHashMap[event.SymbolHash], // get symbol from symbolHashMap with symbol hash as key
				Price:       priceFloat,
				Timestamp:   event.Timestamp.Uint64(),
			})
	}

	return eventsDTO, nil
}

// Helper function to get coingecko prices.
// Takes in coins and looks at their symbols
func GetCoinGeckoPricesForCoins(coins map[string]int64) (map[string]interface{}, error) {

	// create a string of coin symbols so we can get them all in one request
	ids := ""
	index := 1
	for key := range coins {
		ids += key

		if index != len(coins) {
			ids += ","
		}
		index++
	}

	// send a http get request and parse response
	response, err := utils.SendHTTPGetRequest("https://api.coingecko.com/api/v3/simple/price?ids=" + ids + "&vs_currencies=usd")
	if err != nil {
		return nil, err
	}
	jsonResponse, err := utils.ParseJSONResponse(response)
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}