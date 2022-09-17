package server

import (
	"backendtask/blockchain"
	"backendtask/config"
	"math/big"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Endpoint for fetching symbol price history
func CreateSymbolPriceEndpoint(e *echo.Echo, priceSetter *blockchain.Contract) {
	e.GET("/api/price", func(c echo.Context) error {
		symbol := c.QueryParam("symbol")

		if !checkIsSymbolSupported(symbol) {
			return echo.NewHTTPError(http.StatusBadRequest, "Symbol is not monitored by the server")
		}

		priceDTO, err := FetchSymbolDTOPrice(symbol, priceSetter)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed with reason: "+err.Error())
		}

		return c.JSON(http.StatusOK, priceDTO)
	})
}

// Endpoint for fetching symbol price history
func CreateSymbolPriceHistoryEndpoint(e *echo.Echo, priceSetter *blockchain.Contract) {
	e.GET("/api/history", func(c echo.Context) error {
		symbol := c.QueryParam("symbol")

		if !checkIsSymbolSupported(symbol) {
			return echo.NewHTTPError(http.StatusBadRequest, "Symbol is not monitored by the server")
		}

		// check if the timestamp params are valid
		fromTimestamp, errFrom := strconv.ParseInt(c.QueryParam("from"), 10, 64)
		toTimestamp, errTo := strconv.ParseInt(c.QueryParam("to"), 10, 64)
		if errFrom != nil || errTo != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid query values")
		}

		// fetch events from services
		events, err := FetchDTOEvents(
			priceSetter.Address,
			big.NewInt(fromTimestamp),
			big.NewInt(toTimestamp),
			[]string{symbol}, // list of symbols for which to fetch events
		)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch events with reason: "+err.Error())
		}

		return c.JSON(http.StatusOK, events)
	})
}

// checks if the passed symbol is supported
func checkIsSymbolSupported(symbol string) bool {

	foundSymbol := false
	for _, s := range config.Symbols {
		if symbol == s {
			foundSymbol = true
			break
		}
	}
	if !foundSymbol {
		return false
	}
	return true
}
