package main

import (
	"fmt"
	"sort"
	"strconv"
)

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

	fmt.Println(map1)

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

func maps4() {
	fmt.Println("maps4()")
	map1 := map[string]float64{
		"milk":  22.90,
		"bread": 44.50,
	}

	fmt.Println(map1)
	delete(map1, "milk")
	fmt.Println(map1)
}

func maps5() {
	fmt.Println("maps5()")
	map1 := map[string]float64{
		"milk":  22.90,
		"bread": 44.50,
	}

	// not sorted by keys output
	for key, value := range map1 {
		fmt.Println("map1[" + key + "] = " + strconv.FormatFloat(value, 'f', -1, 64))
	}
}

func maps6() {
	fmt.Println("maps6()")
	map1 := map[string]float64{
		"milk":  22.90,
		"bread": 44.50,
	}

	keys := make([]string, 0, len(map1))

	for key := range map1 {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	// sorted by keys output
	for _, key := range keys {
		fmt.Println("map1[" + key + "] = " + strconv.FormatFloat(map1[key], 'f', -1, 64))
	}
}

func maps7() {
	fmt.Println("maps7()")
	price := "$12.36"

	currency := price[0]
	currencySymbol := string(price[0])
	amountAsString := price[1:]

	amount, parseError := strconv.ParseFloat(amountAsString, 64)

	fmt.Println("Original string:", price)
	fmt.Println("Currency (char code):", currency)
	fmt.Println("Currency (char):", currencySymbol)

	if parseError == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse error:", parseError)
	}
}

func maps8() {
	fmt.Println("maps8()")
	price := "€12.36"
	priceAsRune := []rune("€12.36")

	currency := priceAsRune[0]
	currencySymbol := string(priceAsRune[0])
	amountAsString := string(priceAsRune[1:])

	amount, parseError := strconv.ParseFloat(amountAsString, 64)

	fmt.Println("Original string:", price)
	fmt.Println("Currency (char code):", currency)
	fmt.Println("Currency (char code (HEX)):", "0x"+strconv.FormatInt(int64(currency), 16))
	fmt.Println("Currency (char):", currencySymbol)
	fmt.Println("Amount (string)", amountAsString)

	fmt.Println("Length of $12.36", len("$12.36"))
	fmt.Println("Length of €12.36", len("€12.36"))
	fmt.Println("Length of €12.36 (rune)", len(priceAsRune))

	strconv.FormatInt(123, 16)

	if parseError == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse error:", parseError)
	}
}

func maps9() {
	fmt.Println("maps9()")
	price := "€12.36"

	fmt.Println("Original:", price)
	for index, chr := range price {
		fmt.Println("Index:", index, "char code:", chr, "char", string(chr))
	}
}

func maps10() {
	fmt.Println("maps9()")
	price := "€12.36"

	fmt.Println("Original:", price)
	for index, chr := range []byte(price) {
		fmt.Println("Index:", index, "char code:", chr, "char", string(chr))
	}
}
