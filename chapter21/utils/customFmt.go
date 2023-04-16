package utils

import "fmt"

func Printfln(template string, vals ...interface{}) {
	fmt.Printf(template+"\n", vals...)
}
