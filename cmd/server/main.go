package main

import (
	"cashflow/backend/config"
	"cashflow/backend/internal/handlers"
	"cashflow/backend/internal/hooks"
	"cashflow/backend/utils"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

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

	// -------------------------------------------[Migration ]-------------------------------------------
	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	// -------------------------------------------[Custum use]-------------------------------------------
	// [ Currencies ]===========================================
	handlers.BindCurrenciesHooks(app)

	// [ Users ]===========================================
	hooks.BindUsersHooks(app)

	// [ Budgets ]===========================================
	hooks.BindBudgetsHooks(app)

	// [ Transactions ]===========================================
	hooks.BindTransactionsHooks(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
