package utils

import "fmt"

func Printfln(template string, params ...interface{}) {
	fmt.Printf(template+"\n", params...)
}
