package handlers

import (
	"cashflow/backend/utils"
	"net/http"
	"os"

	//"github.com/everapihq/freecurrencyapi-go"
	"github.com/everapihq/currencyapi-go"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
)

// A function to get the currencies data from freecurrencyapi
func GetCurrencies() string {
	currencyapi.Init("cur_live_w6wMNcTij2UewLBsUAkl6cslaoD9LyOtXIYOafTj")
	//freecurrencyapi.Init("fca_live_NneNg9IWsGTcN2B6N34eGSm7ghCxDJwQa4XwHVjS")
	//result := freecurrencyapi.Latest(map[string]string{})
	result := currencyapi.Latest(map[string]string{"base_currency": "USD"})
	return string(result)
}

// Save the currencies data to a file
func SaveCurrencies() {
	currencies := GetCurrencies()
	utils.SaveFile(currencies, "./data/currencies.json")
}

func CurrenciesHandler(c echo.Context) error {
	info := apis.RequestInfo(c)
	admin := info.Admin       // nil if not authenticated as admin
	record := info.AuthRecord // nil if not authenticated as regular auth record

	isLogged := admin != nil || record != nil

	if isLogged {
		dat, err := os.ReadFile("./data/currencies.json")
		if err != nil {
			return err
		}

		/*var jsonMap map[string]interface{}
		err = json.Unmarshal(dat, &jsonMap)
		if err != nil {
			return err
		}*/

		return c.JSON(http.StatusOK, map[string]string{"message": string(dat)})
	} else {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Please login to have access to this endpoint."})
	}
}
