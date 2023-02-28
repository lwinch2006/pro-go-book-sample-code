package models

type Expense interface {
	GetName() string
	GetCost(annual bool) float64
}
