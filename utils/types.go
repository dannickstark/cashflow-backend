package utils

type PaceUnit string

const (
	Day   PaceUnit = "day"
	Week  PaceUnit = "week"
	Month PaceUnit = "month"
	Year  PaceUnit = "year"
)

// -------------------------[ Budgets ]
type TransactionType int64

const (
	Income TransactionType = iota
	Expense
	Transfert
)

type Budget struct {
	Id      string `db:"id" json:"id"`
	Created string `db:"created" json:"created"`

	Amount          string   `db:"amount" json:"amount"`
	Currency        string   `db:"currency" json:"currency"`
	Date            string   `db:"date" json:"date"`
	User            string   `db:"user" json:"user"`
	IncomeCategory  string   `db:"incomeCategory" json:"incomeCategory"`
	ExpenseCategory string   `db:"expenseCategory" json:"expenseCategory"`
	Next            string   `db:"next" json:"next"`
	Repeatable      bool     `db:"repeatable" json:"repeatable"`
	Pace            int64    `db:"pace" json:"pace"`
	PaceUnit        PaceUnit `db:"paceUnit" json:"paceUnit"`
}

// -------------------------[ Budgets ]

type Transaction struct {
	Id      string `db:"id" json:"id"`
	Created string `db:"created" json:"created"`

	Amount          int64    `db:"amount" json:"amount"`
	Fees            int64    `db:"fees" json:"fees"`
	Currency        string   `db:"currency" json:"currency"`
	Date            string   `db:"date" json:"date"`
	User            string   `db:"user" json:"user"`
	IncomeCategory  string   `db:"incomeCategory" json:"incomeCategory"`
	ExpenseCategory string   `db:"expenseCategory" json:"expenseCategory"`
	Next            string   `db:"next" json:"next"`
	Repeatable      bool     `db:"repeatable" json:"repeatable"`
	Pace            int64    `db:"pace" json:"pace"`
	PaceUnit        PaceUnit `db:"paceUnit" json:"paceUnit"`
	Note            string   `db:"note" json:"note"`
	Description     string   `db:"description" json:"description"`
	Picture         string   `db:"picture" json:"picture"`
	Company         string   `db:"company" json:"company"`
	IsExpense       bool     `db:"isExpense" json:"isExpense"`
	IsTransfert     bool     `db:"isTransfert" json:"isTransfert"`
	FromSaving      string   `db:"fromSaving" json:"fromSaving"`
	ToSaving        string   `db:"toSaving" json:"toSaving"`
	Instalment      string   `db:"instalment" json:"instalment"`
	FromAccount     string   `db:"fromAccount" json:"fromAccount"`
	ToAccount       string   `db:"toAccount" json:"toAccount"`
}
