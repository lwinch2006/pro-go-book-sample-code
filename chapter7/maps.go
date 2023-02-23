package main

import "fmt"

func maps1() {
	fmt.Println("maps1()")
	map1 := make(map[string]float64, 10)

	map1["milk"] = 15.36
	map1["bread"] = 34.35

	fmt.Println(map1)
	fmt.Println("Len:", len(map1))
	fmt.Println("Milk price:", map1["milk"])
	fmt.Println("Bread price:", map1["bread"])
}

func maps2() {
	fmt.Println("maps2()")
	map1 := map[string]float64{
		"milk":  22.90,
		"bread": 44.50,
	}

	fmt.Println(map1)
}

func maps3() {
	fmt.Println("maps3()")
	map1 := map[string]float64{
		"milk":  22.90,
		"bread": 44.50,
	}

	if value, keyExists := map1["hat"]; keyExists {
		fmt.Println("Found value[hat]:", value)
	} else {
		fmt.Println("Key has not found: hat")
	}

	if value, keyExists := map1["bread"]; keyExists {
		fmt.Println("Found value[bread]:", value)
	} else {
		fmt.Println("Key has not found: bread")
	}
}
