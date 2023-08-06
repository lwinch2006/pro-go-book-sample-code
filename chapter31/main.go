package main

import (
	"chapter31/services"
	"chapter31/utils"
	"fmt"
)

func main() {
	utils.Printfln("Chapter 31")

	fmt.Println()
	services.UnitTesting1()

	fmt.Println()
	services.UnitTesting2()

	fmt.Println()
	services.UnitTesting3()
}
