package data

import "fmt"

func init() {
	fmt.Println("data.go init function invoked")
}

func GetData() []string {
	return []string{"Milk", "Bread", "Sugar"}
}
