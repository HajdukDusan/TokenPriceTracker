package server

import (
	"backendtask/blockchain"
	"backendtask/config"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Fruit struct {
	Name string `json:"name"`
	Quantity int `json:"quantity"`
  }

func Start(contractAddress string) {
	e := echo.New()

	// Middleware
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	// }))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/history", func(c echo.Context) error {
		symbol := c.QueryParam("symbol")

		foundSymbol := false

		for _, s := range config.Symbols {
			if symbol == s {
				foundSymbol = true
				break;
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

		events, err := blockchain.FetchDTOEvents(
			contractAddress,
			fromTimestamp,
			toTimestamp,
			[]string{symbol},
		)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch events with reason: " + err.Error())
		}

		return c.JSON(http.StatusOK, events)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
