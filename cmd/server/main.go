package main

import (
	"cashflow/backend/config"
	"cashflow/backend/internal/handlers"
	"cashflow/backend/utils"
	"fmt"
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/cron"
)

func main() {
	// load config
	config, err := config.LoadConfig(utils.ConfigPath, utils.ConfigName, utils.ConfigType)
	if err != nil {
		log.Fatalf("could not load configuration file: %v", err)
	}
	fmt.Println(config.Port)

	app := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/currencies", handlers.CurrenciesHandler /* optional middlewares */)
		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		handlers.SaveCurrencies()
		scheduler := cron.New()

		// prints "Hello!" every 2 minutes
		scheduler.MustAdd("getCurrencies", "0 0 * * 0", func() {
			handlers.SaveCurrencies()
		})

		scheduler.Start()

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
