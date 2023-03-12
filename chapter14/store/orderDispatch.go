package store

type DispatchNotification struct {
	Customer string
	*Product
	Quantity int
}
