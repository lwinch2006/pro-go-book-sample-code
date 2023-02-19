package main

import "fmt"

func slice1() {
	fmt.Println("slice1()")

	arr1 := make([]string, 3)
	arr1[0] = "Milk"
	arr1[1] = "Bread"
	arr1[2] = "Beer"

	arr2 := []string{"Milk", "Bread", "Beer"}

	fmt.Println(arr1)
	fmt.Println(arr2)
}

func slice2() {
	fmt.Println("slice2()")
	arr1 := []string{"Milk", "Bread", "Beer"}

	fmt.Println("Len of arr1:", len(arr1))
	fmt.Println("Cap of arr1:", cap(arr1))
	fmt.Println(arr1)

	arr1 = append(arr1, "Wine")

	fmt.Println("Len of arr1:", len(arr1))
	fmt.Println("Cap of arr1:", cap(arr1))
	fmt.Println(arr1)

	arr1 = append(arr1, "Cognac")

	fmt.Println("Len of arr1:", len(arr1))
	fmt.Println("Cap of arr1:", cap(arr1))
	fmt.Println(arr1)

	arr1 = append(arr1, "Whiskey")

	fmt.Println("Len of arr1:", len(arr1))
	fmt.Println("Cap of arr1:", cap(arr1))
	fmt.Println(arr1)

	arr1 = append(arr1, "Vermouth")

	fmt.Println("Len of arr1:", len(arr1))
	fmt.Println("Cap of arr1:", cap(arr1))
	fmt.Println(arr1)
}

func slice3() {
	fmt.Println("slice3()")
	arr1 := []string{"Milk", "Bread", "Beer"}
	arr2 := append(arr1, "Wine", "Tomato")

	arr1[0] = "Kefir"

	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println("Len of arr1:", len(arr1))
	fmt.Println("Cap of arr1:", cap(arr1))

	fmt.Println("Len of arr2:", len(arr2))
	fmt.Println("Cap of arr2:", cap(arr2))
}

func slice4() {
	fmt.Println("slice4()")
	arr1 := make([]string, 3, 6)
	arr1[0] = "Milk"
	arr1[1] = "Bread"
	arr1[2] = "Beer"

	fmt.Println("Length:", len(arr1))
	fmt.Println("Capacity:", cap(arr1))

	tempArr := []string{"Wine", "Tomato"}

	arr2 := append(arr1, "Wine", "Tomato")
	arr3 := append(arr1, tempArr...)

	arr1[0] = "Kefir"

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

func slice5() {
	fmt.Println("slice5()")
	arr1 := [4]string{"Milk", "Bread", "Beer", "Wine"}

	arr2 := arr1[1:3]
	arr3 := arr1[:]

	arr1[1] = "Cake"

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	fmt.Println()

	fmt.Println("Len of arr2:", len(arr2))
	fmt.Println("Cap of arr2:", cap(arr2))

	fmt.Println("Len of arr3:", len(arr3))
	fmt.Println("Cap of arr3:", cap(arr3))

	fmt.Println()

	arr2 = append(arr2, "Cognac")

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

func slice6() {
	fmt.Println("slice6()")
	arr1 := [4]string{"Milk", "Bread", "Beer", "Wine"}

	arr2 := arr1[1:3]
	arr3 := arr1[:]

	arr1[1] = "Cake"

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	fmt.Println()

	fmt.Println("Len of arr2:", len(arr2))
	fmt.Println("Cap of arr2:", cap(arr2))

	fmt.Println("Len of arr3:", len(arr3))
	fmt.Println("Cap of arr3:", cap(arr3))

	fmt.Println()

	arr2 = append(arr2, "Cognac")

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	fmt.Println()

	fmt.Println("Len of arr2:", len(arr2))
	fmt.Println("Cap of arr2:", cap(arr2))

	fmt.Println("Len of arr3:", len(arr3))
	fmt.Println("Cap of arr3:", cap(arr3))

	arr2 = append(arr2, "Tekila")

	fmt.Println()

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	fmt.Println()

	fmt.Println("Len of arr2:", len(arr2))
	fmt.Println("Cap of arr2:", cap(arr2))

	fmt.Println("Len of arr3:", len(arr3))
	fmt.Println("Cap of arr3:", cap(arr3))
}

func slice7() {
	fmt.Println("slice7()")
	arr1 := [4]string{"Milk", "Bread", "Beer", "Wine"}

	//1:3:3 means length = 2 & capacity = 2 so that new underlying array be created when append happens
	//1:3:4 means length = 2 & capacity = 3 so that exising underlying array be used when append happens
	arr2 := arr1[1:3:3]
	arr3 := arr1[:]

	arr1[1] = "Cake"

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	fmt.Println("Len of arr2:", len(arr2))
	fmt.Println("Cap of arr2:", cap(arr2))

	fmt.Println("Len of arr3:", len(arr3))
	fmt.Println("Cap of arr3:", cap(arr3))

	arr2 = append(arr2, "Cognac")

	fmt.Println()

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	fmt.Println("Len of arr2:", len(arr2))
	fmt.Println("Cap of arr2:", cap(arr2))

	fmt.Println("Len of arr3:", len(arr3))
	fmt.Println("Cap of arr3:", cap(arr3))
}

func slice8() {
	fmt.Println("slice8()")
	arr1 := [4]string{"Milk", "Bread", "Beer", "Wine"}

	arr2 := arr1[1:]
	arr3 := arr2[1:3]

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	// Here arr2 and arr3 still has same arr1 as underlying array
	arr2[1] = "Tequila"

	fmt.Println()

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	// Here new underlying array be created for arr2 so that relation between arr2 and arr3 is lost
	arr2 = append(arr2, "Cognac")

	fmt.Println()

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	// This change affects only arr2 as new underlying array been created
	arr2[1] = "Whiskey"

	fmt.Println()

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

func slice9() {
	fmt.Println("slice9()")
	arr1 := [4]string{"Milk", "Bread", "Beer", "Wine"}
	arr2 := arr1[1:]
	arr3 := arr2[1:3]

	arr1[2] = "Tequila"

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	// Copy funcion does not cause creating new underlying array
	copy(arr3, arr2)

	fmt.Println()

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

func slice10() {
	fmt.Println("slice10()")
	arr1 := [4]string{"Milk", "Bread", "Beer", "Wine"}
	arr2 := arr1[1:]
	arr3 := make([]string, 2)

	//arr3[0] = arr2[0]
	//arr3[1] = arr2[1]
	// or
	copy(arr3, arr2)

	arr1[2] = "Tequila"

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	var arr4 []string
	copy(arr4, arr2)

	fmt.Println()
	fmt.Println(arr4) // arr4 is empty since it is not initialized

	copy(arr3[1:], arr2[2:3])

	fmt.Println()
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}
