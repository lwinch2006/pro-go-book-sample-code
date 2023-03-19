package main

import (
	"fmt"
	"regexp"
)

func RegExp1() {
	fmt.Println("RegExp1()")

	description := "Milk is for drinking."

	fmt.Println("Original:", description)
	if match, err := regexp.MatchString("[A-z]ilk", description); err == nil {
		fmt.Println("Matched with pattern [A-z]ilk:", match)
	} else {
		fmt.Println("Error:", err)
	}
}

func RegExp2() {
	fmt.Println("RegExp2()")

	description := "Milk is for drinking."
	pattern, _ := regexp.Compile("[A-z]ilk")

	fmt.Println("Original:", description)
	fmt.Println("Matched with pattern [A-z]ilk:", pattern.MatchString(description))
}

func regexp3_1(s string, indices []int) string {
	return s[indices[0]:indices[1]]
}

func RegExp3() {
	fmt.Println("RegExp3()")

	description := "Milk is for drinking."
	pattern, _ := regexp.Compile("is|for")

	fmt.Println("Original:", description)

	indices1 := pattern.FindStringIndex(description)
	fmt.Println("First string index:", indices1[0], "-", indices1[1], ":", ">>"+regexp3_1(description, indices1)+"<<")

	indices2 := pattern.FindAllStringIndex(description, -1)
	for i, idx := range indices2 {
		fmt.Println("Position:", i, "Indices:", idx[0], "-", idx[1], ":", ">>"+regexp3_1(description, idx)+"<<")
	}
}

func RegExp4() {
	fmt.Println("RegExp4()")

	description := "Milk is for drinking."
	pattern, _ := regexp.Compile(" |is|for")
	slice1 := pattern.Split(description, -1)

	fmt.Println("Original:", description)
	fmt.Println("Splitting by \" \" or \"is\" or \"for\"")
	for _, item := range slice1 {
		if item == "" {
			continue
		}

		fmt.Println(">>" + item + "<<")
	}
}

func RegExp5() {
	fmt.Println("RegExp5()")

	description := "Milk is for drinking."
	pattern, _ := regexp.Compile("[A-z]* is for [A-z]*")

	str := pattern.FindString(description)

	fmt.Println("Original:", description)
	fmt.Println("Match:", str)
}

func RegExp6() {
	fmt.Println("RegExp6()")

	description := "Milk is for drinking."
	pattern, _ := regexp.Compile("([A-z]*) is for ([A-z]*)")
	slice1 := pattern.FindStringSubmatch(description)

	fmt.Println("Original:", description)
	fmt.Println("Found matches:")
	for i, item := range slice1 {
		fmt.Println("Index:", i, "Item:", item)
	}
}

func RegExp7() {
	fmt.Println("RegExp7()")

	description := "Milk is for drinking."
	pattern, _ := regexp.Compile("(?P<object>[A-z]*) is for (?P<action>[A-z]*)")
	slice1 := pattern.FindStringSubmatch(description)

	fmt.Println("Original:", description)
	fmt.Println("Found matches:")
	for _, name := range []string{"object", "action"} {
		fmt.Println("slice1["+name+"]:", slice1[pattern.SubexpIndex(name)])
	}
}

func RegExp8() {
	fmt.Println("RegExp8()")

	description := "Milk is for drinking."
	pattern, _ := regexp.Compile("(?P<object>[A-z]*) is for (?P<action>[A-z]*)")
	template := "Object: ${object}, action: ${action}"

	fmt.Println("Original:", description)
	fmt.Println("Replaced:", pattern.ReplaceAllString(description, template))
}

func regexp9_1(_ string) string {
	return "replacement"
}

func RegExp9() {
	fmt.Println("RegExp9()")

	description := "Milk is for drinking."
	pattern, _ := regexp.Compile("(?P<object>[A-z]*) is for (?P<action>[A-z]*)")

	fmt.Println("Original:", description)
	fmt.Println("Replaced:", pattern.ReplaceAllStringFunc(description, regexp9_1))
}
