package server

import (
	"backendtask/config"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func CreateSymbolPriceHistoryEndpoint(e *echo.Echo, priceSetterAddress string) {
	e.GET("/api/history", func(c echo.Context) error {
		symbol := c.QueryParam("symbol")

		// get more than one symbol?

		foundSymbol := false

		for _, s := range config.Symbols {
			if symbol == s {
				foundSymbol = true
				break
			}
		}

		if !foundSymbol {
			return echo.NewHTTPError(http.StatusBadRequest, "Symbol is not monitored by the server")
		}

		fromTimestamp, errFrom := strconv.ParseInt(c.QueryParam("from"), 10, 64)
		toTimestamp, errTo := strconv.ParseInt(c.QueryParam("to"), 10, 64)
		if errFrom != nil || errTo != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid query values")
		}

		events, err := FetchDTOEvents(
			priceSetterAddress,
			fromTimestamp,
			toTimestamp,
			[]string{symbol},
		)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch events with reason: "+err.Error())
		}

		return c.JSON(http.StatusOK, events)
	})
}
