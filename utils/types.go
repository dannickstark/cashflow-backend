package utils

// -------------------------[ Budgets ]
type BudgetType int64

const (
	Income BudgetType = iota
	Expense
	Transfert
)

type PaceUnit string

const (
	Day   PaceUnit = "day"
	Week  PaceUnit = "week"
	Month PaceUnit = "month"
	Year  PaceUnit = "year"
)

type Budget struct {
	Id              string   `db:"id" json:"id"`
	Amount          string   `db:"amount" json:"amount"`
	Currency        string   `db:"currency" json:"currency"`
	Date            string   `db:"date" json:"date"`
	User            string   `db:"user" json:"user"`
	IncomeCategory  string   `db:"incomeCategory" json:"incomeCategory"`
	ExpenseCategory string   `db:"expenseCategory" json:"expenseCategory"`
	Next            string   `db:"next" json:"next"`
	Pace            int64    `db:"pace" json:"pace"`
	PaceUnit        PaceUnit `db:"paceUnit" json:"paceUnit"`
}
