package models

type Customer struct {
	Name, City string
}

func NewCustomer(name, city string) *Customer {
	return &Customer{name, city}
}

func (c *Customer) GetName() string {
	return c.Name
}

func (c *Customer) TestFunc() string {
	return "Test"
}

func (c *Customer) testFunc2() string {
	return "Test"
}
