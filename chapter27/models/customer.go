package models

type Customer struct {
	Name, City string
}

func NewCustomer(name, city string) *Customer {
	return &Customer{name, city}
}
