package hooks

import (
	"cashflow/backend/utils"
	"log"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/cron"
)

// A function to bind all the hooks for the transactions
func BindTransactionsHooks(app *pocketbase.PocketBase) {
	// Add scheduled jobs
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		scheduler := cron.New()

		// Create a scheduler to create new transaction each day
		scheduler.MustAdd("cloneDailyTransactions", "0 0 1 * *", func() {
			if err := CloneTransactions(app, utils.Day); err != nil {
				log.Println(err.Error())
			}
		})

		// Create a scheduler to create new transaction each week
		scheduler.MustAdd("cloneWeeklyTransactions", "0 0 1 * *", func() {
			if err := CloneTransactions(app, utils.Week); err != nil {
				log.Println(err.Error())
			}
		})

		// Create a scheduler to create new transaction each month
		scheduler.MustAdd("cloneMonthlyTransactions", "0 0 1 * *", func() {
			if err := CloneTransactions(app, utils.Month); err != nil {
				log.Println(err.Error())
			}
		})

		// Create a scheduler to create new transaction each year
		scheduler.MustAdd("cloneAnnualTransactions", "0 0 1 * *", func() {
			if err := CloneTransactions(app, utils.Year); err != nil {
				log.Println(err.Error())
			}
		})

		scheduler.Start()

		return nil
	})

	// Add hooks to execute after lunch of the app
	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
		// Run jobs that will late be run by the scheduler
		if err := CloneTransactions(app, utils.Day); err != nil {
			return err
		}
		if err := CloneTransactions(app, utils.Week); err != nil {
			return err
		}
		if err := CloneTransactions(app, utils.Month); err != nil {
			return err
		}
		if err := CloneTransactions(app, utils.Year); err != nil {
			return err
		}

		return nil
	})
}

func getTransactions(app *pocketbase.PocketBase, category utils.TransactionType, paceUnit utils.PaceUnit) ([]utils.Transaction, error) {
	query := app.Dao().DB().
		Select("transactions.*").
		From("transactions").
		Where(dbx.HashExp{"repeatable": true, "next": ""}).
		AndWhere(dbx.NewExp("paceUnit = {:paceUnit}", dbx.Params{"paceUnit": paceUnit})).
		OrderBy("date ASC").
		Limit(10)

	transactions := []utils.Transaction{}
	if err := query.All(&transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func checkIfTransCorrectPace(transaction utils.Transaction) bool {
	return utils.CheckIfCorrecPace(transaction.Date, int(transaction.Pace), transaction.PaceUnit)
}

// A function to create a new bugget each month
func CreateTransaction(app *pocketbase.PocketBase, transactionData map[string]any) (*models.Record, error) {
	collection, err := app.Dao().FindCollectionByNameOrId("transactions")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	record := models.NewRecord(collection)

	form := forms.NewRecordUpsert(app, record)

	form.LoadData(transactionData)

	if err := form.Submit(); err != nil {
		return nil, err
	}
	return record, nil
}

func computeTransNextDate(transaction utils.Transaction) string {
	return utils.ComputeNextDate(transaction.Date, int(transaction.Pace), transaction.PaceUnit)
}

// A function to clone the transactions
func CloneTransactions(app *pocketbase.PocketBase, paceUnit utils.PaceUnit) error {
	transactions, err := getTransactions(app, utils.Income, paceUnit)
	if err != nil {
		return err
	}
	log.Printf("Found %d transactions", len(transactions))

	// Filter to get transactions with the pace matching to today
	transactions = utils.FilterList(transactions, checkIfTransCorrectPace)

	// Iterate over all transactions
	for _, transaction := range transactions {
		// Clone the transaction
		newData := map[string]any{
			"amount":          transaction.Amount,
			"currency":        transaction.Currency,
			"date":            computeTransNextDate(transaction),
			"user":            transaction.User,
			"incomeCategory":  transaction.IncomeCategory,
			"expenseCategory": transaction.ExpenseCategory,
			"fees":            transaction.Fees,
			"note":            transaction.Note,
			"description":     transaction.Description,
			"picture":         transaction.Picture,
			"company":         transaction.Company,
			"isExpense":       transaction.IsExpense,
			"isTransfert":     transaction.IsTransfert,
			"fromSaving":      transaction.FromSaving,
			"toSaving":        transaction.ToSaving,
			"instalment":      transaction.Instalment,
			"fromAccount":     transaction.FromAccount,
			"toAccount":       transaction.ToAccount,
			"repeatable":      transaction.Repeatable,
			"pace":            transaction.Pace,
			"paceUnit":        transaction.PaceUnit,
		}
		// Create the transaction
		newTransaction, err := CreateTransaction(app, newData)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		// Update the current transaction
		oldTransaction, err := app.Dao().FindRecordById("transactions", transaction.Id)
		if err != nil {
			return err
		}
		oldTransaction.Set("next", newTransaction.Get("id"))
		if err := app.Dao().SaveRecord(oldTransaction); err != nil {
			return err
		}
	}

	return nil
}
