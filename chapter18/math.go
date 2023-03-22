package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(1)
}

func Math1() {
	fmt.Println("Math1()")

	val1 := 136.816
	val2 := 2.00

	Printfln("Abs: %v", math.Abs(val1))
	Printfln("Ceil: %v", math.Ceil(val1))
	Printfln("Copysign: %v", math.Copysign(val1, -1*val2))
	Printfln("Floor: %v", math.Floor(val1))
	Printfln("Max: %v", math.Max(val1, val2))
	Printfln("Min: %v", math.Min(val1, val2))
	Printfln("Mod: %v", math.Mod(val1, val2))
	Printfln("Pow: %v", math.Pow(val1, val2))
	Printfln("Round: %v", math.Round(val1))
	Printfln("RoundToEven: %v", math.RoundToEven(val1))
}

func Math2() {
	fmt.Println("Math2()")

	for i := 0; i < 5; i++ {
		Printfln("Random value: %v", rand.Int())
	}
}

func Math3() {
	fmt.Println("Math3()")

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		Printfln("Random value: %v", rand.Int())
	}
}

func Math4() {
	fmt.Println("Math4()")

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		Printfln("Random value: %v", rand.Intn(10))
	}
}

func math5_1(min, max int) int {
	return rand.Intn(max-min) + min
}

func Math5() {
	fmt.Println("Math5()")

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		Printfln("Random value: %v", math5_1(10, 20))
	}
}

func Math6() {
	fmt.Println("Math6()")

	rand.Seed(time.Now().UnixNano())

	names := []string{"Alice", "Bob", "Charlie", "Dave", "Ester"}

	rand.Shuffle(len(names), func(i, j int) {
		names[i], names[j] = names[j], names[i]
	})

	for i := 0; i < len(names); i++ {
		Printfln("Name: %v", names[i])
	}
}
