package models

type CategorySubtotalMessage struct {
	Category      string
	Subtotal      float64
	TerminalError interface{}
}
