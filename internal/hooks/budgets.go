package hooks

import (
	"cashflow/backend/utils"
	"fmt"
	"log"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/cron"
)

// A function to bind all the hooks for the budgets
func BindBudgetsHooks(app *pocketbase.PocketBase) {
	// Add scheduled jobs
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		scheduler := cron.New()

		// Create a scheduler to get the create new budget each day
		scheduler.MustAdd("cloneDailyBudgets", "0 0 1 * *", func() {
			if err := CloneBudgets(app, utils.Day); err != nil {
				log.Println(err.Error())
			}
		})

		// Create a scheduler to get the create new budget each week
		scheduler.MustAdd("cloneWeeklyBudgets", "0 0 1 * *", func() {
			if err := CloneBudgets(app, utils.Week); err != nil {
				log.Println(err.Error())
			}
		})

		// Create a scheduler to get the create new budget each month
		scheduler.MustAdd("cloneMonthlyBudgets", "0 0 1 * *", func() {
			if err := CloneBudgets(app, utils.Month); err != nil {
				log.Println(err.Error())
			}
		})

		// Create a scheduler to get the create new budget each year
		scheduler.MustAdd("cloneAnnualBudgets", "0 0 1 * *", func() {
			if err := CloneBudgets(app, utils.Year); err != nil {
				log.Println(err.Error())
			}
		})

		scheduler.Start()

		return nil
	})

	// Add hooks to execute after lunch of the app
	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
		// Run jobs that will late be run by the scheduler
		if err := CloneBudgets(app, utils.Day); err != nil {
			return err
		}
		if err := CloneBudgets(app, utils.Week); err != nil {
			return err
		}
		if err := CloneBudgets(app, utils.Month); err != nil {
			return err
		}
		if err := CloneBudgets(app, utils.Year); err != nil {
			return err
		}

		return nil
	})
}

// A function to generate the date condition depending of the paceUnit
/* func getDateCondition(dateKey string, paceUnit utils.PaceUnit) string {
	dateCondition := ""

	switch paceUnit {
	case utils.Day:
		dateCondition = fmt.Sprintf("%s BETWEEN '%s' AND '%s'",
			dateKey,
			carbon.Yesterday().StartOfDay().ToDateTimeString(),
			carbon.Yesterday().EndOfDay().ToDateTimeString(),
		)

	case utils.Week:
		dateCondition = fmt.Sprintf("%s BETWEEN '%s' AND '%s'",
			dateKey,
			carbon.Now().SubWeek().StartOfWeek().ToDateTimeString(),
			carbon.Now().SubWeek().EndOfWeek().ToDateTimeString(),
		)

	case utils.Month:
		dateCondition = fmt.Sprintf("%s BETWEEN '%s' AND '%s'",
			dateKey,
			carbon.Now().SubMonth().StartOfMonth().ToDateTimeString(),
			carbon.Now().SubMonth().EndOfMonth().ToDateTimeString(),
		)

	case utils.Year:
		dateCondition = fmt.Sprintf("%s BETWEEN '%s' AND '%s'",
			dateKey,
			carbon.Now().SubYear().StartOfYear().ToDateTimeString(),
			carbon.Now().SubYear().EndOfYear().ToDateTimeString(),
		)
	}

	return dateCondition
} */

func checkIfBudgetCorrectPace(budget utils.Budget) bool {
	return utils.CheckIfCorrecPace(budget.Date, int(budget.Pace), budget.PaceUnit)
}

func computeBudgNextDate(budget utils.Budget) string {
	return utils.ComputeNextDate(budget.Date, int(budget.Pace), budget.PaceUnit)
}

// A function to create a new bugget each month
func CreateBudget(app *pocketbase.PocketBase, budgetData map[string]any) (*models.Record, error) {
	collection, err := app.Dao().FindCollectionByNameOrId("budgets")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	record := models.NewRecord(collection)

	form := forms.NewRecordUpsert(app, record)

	form.LoadData(budgetData)

	if err := form.Submit(); err != nil {
		return nil, err
	}
	return record, nil
}

func getBudgetByCat(app *pocketbase.PocketBase, category utils.TransactionType, paceUnit utils.PaceUnit) ([]utils.Budget, error) {
	budgets := []utils.Budget{}
	filter_cat := ""

	switch category {
	case utils.Income:
		filter_cat = "incomeCategory"

	case utils.Expense:
		filter_cat = "expenseCategory"
	}

	err := app.Dao().DB().
		NewQuery(fmt.Sprintf(`
			WITH numbered_budgets AS (
				SELECT *, ROW_NUMBER() OVER(PARTITION BY %s ORDER BY date DESC) AS row_number
				FROM budgets
				WHERE %s != ''
					AND repeatable = true
					AND paceUnit = {:paceUnit}
					AND next = ''
			)
			SELECT *
			FROM numbered_budgets
			WHERE row_number = 1;
		`, filter_cat, filter_cat)).
		Bind(dbx.Params{
			"paceUnit": paceUnit,
		}).
		All(&budgets)

	if err != nil {
		return nil, err
	}

	return budgets, nil
}

// A function to clone the budgets of last month
func CloneBudgets(app *pocketbase.PocketBase, paceUnit utils.PaceUnit) error {
	incomeBudgets, err := getBudgetByCat(app, utils.Income, paceUnit)
	if err != nil {
		return err
	}
	log.Printf("Found %d incomeBudgets", len(incomeBudgets))

	expensesBudgets, err := getBudgetByCat(app, utils.Expense, paceUnit)
	if err != nil {
		return err
	}
	log.Printf("Found %d expensesBudgets", len(expensesBudgets))

	budgets := append(incomeBudgets, expensesBudgets...)
	log.Printf("Found %d budgets", len(budgets))

	// Filter to get budgets with the pace matching to today
	budgets = utils.FilterList(budgets, checkIfBudgetCorrectPace)

	// Iterate over all budgets
	for _, budget := range budgets {
		// Clone the budget
		newData := map[string]any{
			"amount":          budget.Amount,
			"currency":        budget.Currency,
			"date":            computeBudgNextDate(budget),
			"user":            budget.User,
			"incomeCategory":  budget.IncomeCategory,
			"expenseCategory": budget.ExpenseCategory,
			"repeatable":      budget.Repeatable,
			"pace":            budget.Pace,
			"paceUnit":        budget.PaceUnit,
		}
		// Create the budget
		newBudget, err := CreateBudget(app, newData)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		// Update the current budget
		oldBudget, err := app.Dao().FindRecordById("budgets", budget.Id)
		if err != nil {
			return err
		}
		oldBudget.Set("next", newBudget.Get("id"))
		if err := app.Dao().SaveRecord(oldBudget); err != nil {
			return err
		}
	}

	return nil
}
