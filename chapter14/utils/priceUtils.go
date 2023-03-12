package utils

import (
	"strconv"
)

func ToCurrency(price float64) string {
	return "$" + strconv.FormatFloat(price, 'f', 2, 64)
}
