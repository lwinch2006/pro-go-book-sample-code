package main

import "fmt"

func StringsScanning1() {
	fmt.Println("StringsScanning1()")

	var name, category string
	var price float64

	fmt.Println("Enter new product:")

	// Returns error if not pointers used
	n, err := fmt.Scan(&name, &category, &price)

	if err == nil {
		fmt.Printf("Number of values read: %v", n)
		fmt.Println()
		fmt.Printf("New product \"%v\" of category \"%v\" at price $%.2f added", name, category, price)
	} else {
		fmt.Printf("Error received: %v", err.Error())
	}

	fmt.Println()
}

func stringsScanning2_1(items ...string) (itemsP []interface{}) {
	itemsP = make([]interface{}, len(items))
	for i := 0; i < len(items); i++ {
		itemsP[i] = &items[i]
	}

	return
}

func StringsScanning2() {
	fmt.Println("StringsScanning2()")

	vals := make([]string, 3)
	valsP := stringsScanning2_1(vals...)

	fmt.Println("Enter 3 string values:")

	n, err := fmt.Scan(valsP...)

	if err == nil {
		fmt.Printf("Number of values read: %v", n)
		fmt.Println()
		fmt.Printf("Values: %v", vals)
	} else {
		fmt.Printf("Error received: %v", err.Error())
	}

	fmt.Println()
}

func StringsScanning3() {
	fmt.Println("StringsScanning3()")

	vals := make([]string, 3)
	valsP := stringsScanning2_1(vals...)

	fmt.Println("Enter 3 string values:")

	n, err := fmt.Scanln(valsP...)

	if err == nil {
		fmt.Printf("Number of values read: %v", n)
		fmt.Println()
		fmt.Printf("Values: %v", vals)
	} else {
		fmt.Printf("Error received: %v", err.Error())
	}

	fmt.Println()
}

func StringsScanning4() {
	fmt.Println("StringsScanning4()")

	source := "test1 test2 test3 test4"
	items := make([]string, 3)
	itemsP := stringsScanning2_1(items...)

	n, err := fmt.Sscan(source, itemsP...)

	if err == nil {
		fmt.Printf("Number of values read: %v", n)
		fmt.Println()
		fmt.Printf("Values: %v", items)
	} else {
		fmt.Printf("Error received: %v", err.Error())
	}

	fmt.Println()
}

func StringsScanning5() {
	fmt.Println("StringsScanning5()")

	source := "test1 test2 test3 test4\n"
	items := make([]string, 3)
	itemsP := stringsScanning2_1(items...)

	n, err := fmt.Sscanln(source, itemsP...)

	if err == nil {
		fmt.Printf("Number of values read: %v", n)
		fmt.Println()
		fmt.Printf("Values: %v", items)
	} else {
		fmt.Printf("Error received: %v", err.Error())
	}

	fmt.Println()
}

func StringsScanning6() {
	fmt.Println("StringsScanning6()")

	source := "Given values: test1 test2 test3 test4"
	template := "Given values: %s %s %s"
	items := make([]string, 3)
	itemsP := stringsScanning2_1(items...)

	n, err := fmt.Sscanf(source, template, itemsP...)

	if err == nil {
		fmt.Printf("Number of values read: %v", n)
		fmt.Println()
		fmt.Printf("Values: %v", items)
	} else {
		fmt.Printf("Error received: %v", err.Error())
	}

	fmt.Println()
}
