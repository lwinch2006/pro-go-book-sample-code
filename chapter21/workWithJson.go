package main

import (
	"chapter21/models"
	"chapter21/utils"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func WorkWithJson1() {
	fmt.Println("WorkWithJson1()")

	var w1 strings.Builder
	val := 111
	ptr := &val

	enc := json.NewEncoder(&w1)

	for _, value := range []interface{}{true, "Milk", 1, 23.36, 'Æ', ptr} {
		enc.Encode(value)
	}

	fmt.Println(w1.String())
}

func WorkWithJson2() {
	fmt.Println("WorkWithJson2()")

	arr1 := []string{"Milk", "Bread", "Beer"}
	arr2 := []int{111, 222, 333}
	var arr3 [4]byte
	copy(arr3[0:], arr1[0])
	arr4 := []byte(arr1[0])

	var w1 strings.Builder
	enc := json.NewEncoder(&w1)

	for _, value := range []interface{}{arr1, arr2, arr3, arr4} {
		enc.Encode(value)
	}

	fmt.Println(w1.String())
}

func WorkWithJson3() {
	fmt.Println("WorkWithJson3()")

	map1 := map[string]float64{
		"Milk":  23.34,
		"Bread": 45.56,
	}

	var w1 strings.Builder
	enc := json.NewEncoder(&w1)
	enc.Encode(map1)

	fmt.Println(w1.String())
}

func WorkWithJson4() {
	fmt.Println("WorkWithJson4()")

	p1 := models.NewProduct("Milk", "Food", 23.34)

	var w1 strings.Builder
	enc := json.NewEncoder(&w1)
	enc.Encode(p1)

	fmt.Println(w1.String())
}

func WorkWithJson5() {
	fmt.Println("WorkWithJson5()")

	p1 := models.NewDiscountedProduct("Milk", "Food", 23.34, 5.00)

	var w1 strings.Builder
	enc := json.NewEncoder(&w1)
	enc.Encode(p1)

	fmt.Println(w1.String())
}

func WorkWithJson6() {
	fmt.Println("WorkWithJson6()")

	p1 := models.DiscountedProduct{
		Discount: 12.23,
	}

	var w1 strings.Builder
	enc := json.NewEncoder(&w1)
	enc.Encode(p1)

	fmt.Println(w1.String())
}

func WorkWithJson7() {
	fmt.Println("WorkWithJson7()")

	p1 := models.NewDiscountedProduct("Milk", "Food", 23.36, 5.43)
	p2 := models.NewPerson("John")
	arr1 := []models.Named{p1, p2}
	var w1 strings.Builder
	enc := json.NewEncoder(&w1)
	enc.Encode(arr1)

	fmt.Println(w1.String())
}

func WorkWithJson8() {
	fmt.Println("WorkWithJson8()")

	r1 := strings.NewReader(`true "Hello" 99.99 200`)
	dec := json.NewDecoder(r1)

	values := []interface{}{}

	for {
		var value interface{}
		err := dec.Decode(&value)

		if err != nil {
			if err != io.EOF {
				utils.Printfln("Error: %v", err.Error())
			}

			break
		}

		values = append(values, value)
	}

	for _, valueItem := range values {
		utils.Printfln("Value (%T): %v", valueItem, valueItem)
	}
}

func WorkWithJson9() {
	fmt.Println("WorkWithJson9()")

	r1 := strings.NewReader(`true "Hello" 99.99 200`)
	dec := json.NewDecoder(r1)
	dec.UseNumber()

	values := []interface{}{}

	for {
		var value interface{}
		err := dec.Decode(&value)

		if err != nil {
			if err != io.EOF {
				utils.Printfln("Error: %v", err.Error())
			}

			break
		}

		values = append(values, value)
	}

	for _, valueItem := range values {
		if num, ok := valueItem.(json.Number); ok {
			if ival, err := num.Int64(); err == nil {
				utils.Printfln("Value (int64): %v", ival)
			} else if fval, err := num.Float64(); err == nil {
				utils.Printfln("Value (float64): %v", fval)
			} else {
				utils.Printfln("Value (string): %v", num.String())
			}
		} else {
			utils.Printfln("Value (%T): %v", valueItem, valueItem)
		}
	}
}

func WorkWithJson10() {
	fmt.Println("WorkWithJson10()")

	r1 := strings.NewReader(`true "Hello" 99.99 200`)
	dec := json.NewDecoder(r1)

	var value1 bool
	var value2 string
	var value3 float64
	var value4 int

	values := []interface{}{&value1, &value2, &value3, &value4}

	for _, valueItem := range values {
		err := dec.Decode(valueItem)

		if err != nil {
			utils.Printfln("Error: %v", err.Error())
		}
	}

	utils.Printfln("Value1 (%T): %v", value1, value1)
	utils.Printfln("Value2 (%T): %v", value2, value2)
	utils.Printfln("Value3 (%T): %v", value3, value3)
	utils.Printfln("Value4 (%T): %v", value4, value4)
}

func WorkWithJson11() {
	fmt.Println("WorkWithJson11()")

	r1 := strings.NewReader(`[10,20,30]["Milk","Food",23.36]`)
	dec := json.NewDecoder(r1)

	values := []interface{}{}

	for {
		var value interface{}
		err := dec.Decode(&value)

		if err != nil {
			if err != io.EOF {
				utils.Printfln("Error: %v", err.Error())
			}

			break
		}

		values = append(values, value)
	}

	for _, valueItem := range values {
		utils.Printfln("Value (%T): %v", valueItem, valueItem)
	}
}

func WorkWithJson12() {
	fmt.Println("WorkWithJson12()")

	r1 := strings.NewReader(`[10,20,30]["Milk","Food",23.36]`)
	dec := json.NewDecoder(r1)

	arr1 := []int{}
	arr2 := []interface{}{}

	values := []interface{}{&arr1, &arr2}

	for _, valueItem := range values {
		err := dec.Decode(valueItem)

		if err != nil {
			utils.Printfln("Error: %v", err.Error())
		}
	}

	utils.Printfln("Arr1 (%T): %v", arr1, arr1)
	utils.Printfln("Arr2 (%T): %v", arr2, arr2)
}

func WorkWithJson13() {
	fmt.Println("WorkWithJson13()")

	r1 := strings.NewReader(`{"Milk": 23.36, "Bread": 45.56}`)
	dec := json.NewDecoder(r1)
	m1 := map[string]interface{}{}
	err := dec.Decode(&m1)

	if err != nil {
		utils.Printfln("Error: %v", err.Error())
	} else {
		utils.Printfln("Map (%T): %v", m1, m1)

		for k, v := range m1 {
			utils.Printfln("map[%v]: %v", k, v)
		}
	}
}

func WorkWithJson14() {
	fmt.Println("WorkWithJson14()")

	r1 := strings.NewReader(`
	{"Name": "Milk", "Category": "Food", "Price": 23.34}
	{"name": "Bread", "Category": "Food", "Price": 45.56}
	{"Name": "Beer", "Category": "Food", "Price": 78.89, "InStock": true}
	`)

	dec := json.NewDecoder(r1)

	for {
		var p models.Product
		err := dec.Decode(&p)
		if err != nil {
			if err != io.EOF {
				utils.Printfln("Error: %v", err.Error())
			}
			break
		}
		utils.Printfln("Product %v of category %v with price %v", p.Name, p.Category, p.Price)
	}
}

func WorkWithJson15() {
	fmt.Println("WorkWithJson15()")

	r1 := strings.NewReader(`
	{"Name": "Milk", "Category": "Food", "Price": 23.34}
	{"name": "Bread", "Category": "Food", "Price": 45.56}
	{"Name": "Beer", "Category": "Food", "Price": 78.89, "InStock": true}
	`)

	dec := json.NewDecoder(r1)
	dec.DisallowUnknownFields()

	for {
		var p models.Product
		err := dec.Decode(&p)
		if err != nil {
			if err != io.EOF {
				utils.Printfln("Error: %v", err.Error())
			}
			break
		}
		utils.Printfln("Product %v of category %v with price %v", p.Name, p.Category, p.Price)
	}
}

func WorkWithJson16() {
	fmt.Println("WorkWithJson16()")

	r1 := strings.NewReader(`
	{"Name": "Milk", "Category": "Food", "Price": 23.34, "Offer": "12.23"}
	`)

	dec := json.NewDecoder(r1)
	dec.DisallowUnknownFields()

	for {
		var p models.DiscountedProduct
		err := dec.Decode(&p)
		if err != nil {
			if err != io.EOF {
				utils.Printfln("Error: %v", err.Error())
			}
			break
		}
		utils.Printfln("Product %v of category %v with price %v and discount %v", p.Name, p.Category, p.Price, p.Discount)
	}
}
