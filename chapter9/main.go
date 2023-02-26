package main

import "fmt"

func main() {
	fmt.Println("Chapter 9")

	fmt.Println()
	funcTypes1()

	printFunc := fmt.Println
	_, _ = printFunc()
	funcTypes2("Function 1 being used as variable and passed to funcTypes2() function", printFunc)
	funcTypes2("Function 2 being used as variable and passed to funcTypes2() function", func(a ...any) (n int, err error) {
		n, err = fmt.Println(a...)
		return
	})

	fmt.Println()
	funcTypes3()

	fmt.Println()
	funcTypes4()

	fmt.Println()
	funcTypes5()

	fmt.Println()
	funcTypes6()

	fmt.Println()
	funcTypes7()

	fmt.Println()
	funcTypes8()

	fmt.Println()
	funcTypes9()

	fmt.Println()
	funcTypes10()
}
