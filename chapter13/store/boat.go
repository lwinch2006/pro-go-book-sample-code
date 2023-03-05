package store

type Boat struct {
	*Product
	Capacity  int
	Motorized bool
}

func NewBoat(name, category string, price float64, capacity int, motorized bool) *Boat {
	return &Boat{NewProduct(name, category, price), capacity, motorized}
}
