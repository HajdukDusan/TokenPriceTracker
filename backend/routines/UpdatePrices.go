package routines

import (
	"backendtask/blockchain"
	"backendtask/config"
	"backendtask/utils"
	"fmt"
	"math/big"
	"time"
)

// Goroutine for updating the prices in PriceSetter contract
func UpdatePrices(coins map[string]int64) {

	for true {

		newPrices := getCoinGeckoPricesForCoins(coins)

		updateCoinPrices(coins, newPrices)

		updateContractCoinPrices(coins)

		// stop thread for config.UpdateInterval seconds
		time.Sleep(time.Duration(config.UpdateInterval) * time.Second)
	}
}

// Helper function to update contract coin prices.
// Sends txs as a block, which overrides pending txs when called next time.
func updateContractCoinPrices(coins map[string]int64) {

	// Init the block of txs
	SendTx := blockchain.CreateTxBlock()

	for symbol, price := range coins {

		// get symbol price from contract
		contractPrice, err := blockchain.PriceSetterContract.GetSymbolPrice(symbol)
		if err != nil {
			fmt.Println(err)
		}

		absDif := utils.GetAbs(big.NewInt(0).Sub(contractPrice, big.NewInt(price)))
		minDif := big.NewInt(2)

		// check if new price difference is more than 2%
		if absDif.Mul(absDif, big.NewInt(100)).Cmp(minDif.Mul(contractPrice, minDif)) == 1 {

			// send tx to mempool
			tx, err := SendTx(symbol, big.NewInt(price))
			if err != nil {
				fmt.Println("Failed to send tx!")
				fmt.Println(err)
			} else {
				fmt.Printf("Tx successfully sent: %s\n", tx.Hash().Hex())
			}
		}
	}
}

// Helper function to update coin prices.
// Takes in coins to update and their new prices.
func updateCoinPrices(coins map[string]int64, newPrices map[string]interface{}) {

	fmt.Println("\nNew Prices:")

	for key := range coins {
		coin := newPrices[key].(map[string]interface{})
		price := coin["usd"].(float64)

		fmt.Printf("\t" + key + " -> %f USD\n", price)

		// scale price to int, dont multiply because its a float value
		result, err := utils.ScaleFloatToInt(price)
		if err != nil {
			fmt.Println(err)
			continue
		}

		coins[key] = result
	}
}

// Helper function to get coingecko prices.
// Takes in coins and looks at their symbols
func getCoinGeckoPricesForCoins(coins map[string]int64) map[string]interface{} {

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
		fmt.Println(err)
	}
	jsonResponse, err := utils.ParseJSONResponse(response)
	if err != nil {
		fmt.Println(err)
	}

	return jsonResponse
}