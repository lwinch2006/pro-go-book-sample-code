package main

import "fmt"

func arrays1() {
	arr1 := []string{"Milk", "Bread", "Beer"}    // dynamic size
	arr2 := [...]string{"Milk", "Bread", "Beer"} // fixed size

	fmt.Println(arr1)
	fmt.Println(arr2)

	arr1 = append(arr1, "Wine")
	fmt.Println(arr1)
}

func arrays2() {
	//arr1 := [3]string{"Milk", "Bread", "Beer"}
	//var arr2 [4]string = arr1
}

func arrays3() {
	arr1 := [3]string{"Milk", "Bread", "Beer"}
	arr2 := arr1

	arr1[0] = "Wine"

	fmt.Println(arr1)
	fmt.Println(arr2)
}

func arrays4() {
	arr1 := [3]string{"Milk", "Bread", "Beer"}
	arr2 := &arr1

	arr1[0] = "Wine"

	fmt.Println(arr1)
	fmt.Println(*arr2)
}

func arrays5() {
	//arr1 := []string{"Milk", "Bread", "Beer"}    // dynamic size
	arr2 := [...]string{"Milk", "Bread", "Beer"} // fixed size
	arr3 := [3]string{"Milk", "Bread", "Beer"}   // fixed size

	fmt.Println("arr2 == arr3:", arr2 == arr3)
	//fmt.Println("arr1 == arr2:", arr1 == arr2) // not equal as type mismatched
}
