package handlers

import (
	"cashflow/backend/utils"
	"net/http"
	"os"

	//"github.com/everapihq/freecurrencyapi-go"
	"github.com/everapihq/currencyapi-go"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/cron"
)

// A function to bind all the hooks for the currencies
func BindCurrenciesHooks(app core.App) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/currencies", CurrenciesHandler /* optional middlewares */)
		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// Get the currencies values from the API
		SaveCurrencies()

		// Create a scheduler to get the currencie's data each sunday
		scheduler := cron.New()
		scheduler.MustAdd("getCurrencies", "0 0 * * 0", func() {
			SaveCurrencies()
		})

		scheduler.Start()

		return nil
	})
}

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
	utils.SaveFile(currencies, utils.GetAbsolutePath("data/currencies.json"))
}

func CurrenciesHandler(c echo.Context) error {
	info := apis.RequestInfo(c)
	admin := info.Admin       // nil if not authenticated as admin
	record := info.AuthRecord // nil if not authenticated as regular auth record

	isLogged := admin != nil || record != nil

	if isLogged {
		dat, err := os.ReadFile(utils.GetAbsolutePath("data/currencies.json"))
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
