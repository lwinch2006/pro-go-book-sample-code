package models

type Account struct {
	Number   int
	Expenses []Expense
}

func NewAccount(number int) *Account {
	return &Account{number, make([]Expense, 0)}
}

func (account *Account) AddExpense(expense Expense) {
	account.Expenses = append(account.Expenses, expense)
}
