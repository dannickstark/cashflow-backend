package main

import (
	"cashflow/backend/config"
	"cashflow/backend/internal/handlers"
	"cashflow/backend/utils"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/pocketbase/pocketbase/tools/cron"

	// uncomment once you have at least one .go migration file in the "migrations" directory
	_ "cashflow/backend/migrations"
)

// Start on http://127.0.0.1:8090
func main() {
	// load config
	config, err := config.LoadConfig(utils.ConfigPath, utils.ConfigName, utils.ConfigType)
	if err != nil {
		log.Fatalf("could not load configuration file: %v", err)
	}
	fmt.Println(config.Port)

	app := pocketbase.New()

	// -------------------------------------------[Custum use]-------------------------------------------

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

		scheduler.MustAdd("getCurrencies", "0 0 * * 0", func() {
			handlers.SaveCurrencies()
		})

		scheduler.Start()

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

	// -------------------------------------------[Migration]-------------------------------------------
	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
