package models

type ProductChannelMessage struct {
	Category string
	Subtotal float64
	*CategoryNotExistError
}
