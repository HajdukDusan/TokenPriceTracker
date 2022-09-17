package routines

import (
	"backendtask/blockchain"
	"backendtask/config"
	"backendtask/server"
	"backendtask/utils"
	"fmt"
	"math/big"
	"time"
)

// Goroutine for updating the prices in PriceSetter contract
func UpdatePrices(coins map[string]int64, priceSetter *blockchain.Contract, privKey string) {

	for true {

		newPrices, err := server.GetCoinGeckoPricesForCoins(coins)
		if err != nil {
			fmt.Println(err)
		}

		updateCoinPrices(coins, newPrices)

		updateContractCoinPrices(coins, priceSetter, privKey)

		// stop thread for config.UpdateInterval seconds
		time.Sleep(time.Duration(config.UpdateInterval) * time.Second)
	}
}

// Helper function to update contract coin prices.
// Sends txs as a block, which overrides pending txs when called next time.
func updateContractCoinPrices(coins map[string]int64, priceSetter *blockchain.Contract, privKey string) {

	// Init the block of txs
	SendTx := blockchain.CreateTxBlock(privKey)

	for symbol, price := range coins {

		// get symbol price from contract
		contractPrice, err := priceSetter.GetSymbolPrice(symbol)
		if err != nil {
			fmt.Println(err)
		}

		absDif := utils.GetAbs(big.NewInt(0).Sub(contractPrice, big.NewInt(price)))
		minDif := big.NewInt(2)

		// check if new price difference is more than 2%
		if absDif.Mul(absDif, big.NewInt(100)).Cmp(minDif.Mul(contractPrice, minDif)) == 1 {

			fmt.Println("Updating " + symbol + " price...")

			// send tx to mempool
			tx, err := SendTx(priceSetter, symbol, big.NewInt(price))
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