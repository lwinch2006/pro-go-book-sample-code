package fmt

import "strconv"

func ToCurrency(value float64) string {
	return "$" + strconv.FormatFloat(value, 'f', 2, 64)
}
