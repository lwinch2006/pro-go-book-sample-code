package main

import (
	"chapter17/models"
	"fmt"
)

func StringsFormatting1() {
	fmt.Println("StringsFormatting1()")

	fmt.Println("Product:", models.Milk.Name, "Price:", models.Milk.Price)
	fmt.Print("Product:", models.Milk.Name, "Price:", models.Milk.Price)
	fmt.Println()
}

func StringsFormatting2() {
	fmt.Println("StringsFormatting2()")

	fmt.Printf("Product \"%v\" at price $%4.2f", models.Milk.Name, models.Milk.Price)
	fmt.Println()
}

func stringsFormatting3_1(index int) string {
	if index < len(models.Products) {
		return fmt.Sprintf("Product \"%v\" at price $%4.2f", models.Products[index].Name, models.Products[index].Price)
	}

	return fmt.Errorf("error for index %v", index).Error()
}

func StringsFormatting3() {
	fmt.Println("StringsFormatting3()")

	fmt.Println(stringsFormatting3_1(1))
	fmt.Println(stringsFormatting3_1(10))
}

func StringsFormatting4() {
	fmt.Println("StringsFormatting3()")

	fmt.Printf("Product (v): %v", *models.Milk)
	fmt.Println()
	fmt.Printf("Product (+v): %+v", *models.Milk)
	fmt.Println()
	fmt.Printf("Product (#v): %#v", *models.Milk)
	fmt.Println()
	fmt.Printf("Product (T): %T", *models.Milk)
	fmt.Println()
}

func StringsFormatting5() {
	fmt.Println("StringsFormatting5()")

	fmt.Println(models.Milk)

	fmt.Printf("Product (v): %v", models.Milk)
	fmt.Println()

	fmt.Printf("Product (v): %+v", models.Milk)
	fmt.Println()

	// Since String() extension function defined for pointer here it will be used default printing
	fmt.Printf("Product (v): %v", *models.Milk)
	fmt.Println()
}

func StringsFormatting6() {
	fmt.Println("StringsFormatting6()")

	priceInt := int(models.Milk.Price)

	fmt.Printf("Product \"%v\" at price $%4.2f (float)", models.Milk.Name, models.Milk.Price)
	fmt.Println()

	fmt.Printf("Product \"%v\" at price $%d (decimal)", models.Milk.Name, priceInt)
	fmt.Println()

	fmt.Printf("Product \"%v\" at price $0b%b (binary)", models.Milk.Name, priceInt)
	fmt.Println()

	fmt.Printf("Product \"%v\" at price $%O (octal)", models.Milk.Name, priceInt)
	fmt.Println()

	fmt.Printf("Product \"%v\" at price $0x%X (hex)", models.Milk.Name, priceInt)
	fmt.Println()
}

func StringsFormatting7() {
	fmt.Println("StringsFormatting7()")

	fmt.Printf("Product \"%v\" at price $%4.2f (float)", models.Milk.Name, models.Milk.Price)
	fmt.Println()

	fmt.Printf("Product \"%v\" at price $%b (exponential without decimal)", models.Milk.Name, models.Milk.Price)
	fmt.Println()

	fmt.Printf("Product \"%v\" at price $%E (exponential with decimal)", models.Milk.Name, models.Milk.Price)
	fmt.Println()

	fmt.Printf("Product \"%v\" at price $%F (float)", models.Milk.Name, models.Milk.Price)
	fmt.Println()

	fmt.Printf("Product \"%v\" at price $%G (g/G)", models.Milk.Name, models.Milk.Price)
	fmt.Println()

	fmt.Printf("Product \"%v\" at price $%X (hex)", models.Milk.Name, models.Milk.Price)
	fmt.Println()
}

func StringsFormatting8() {
	fmt.Println("StringsFormatting8()")

	fmt.Printf("Product \"%v\" at price >>$%4.2f<<", models.Milk.Name, models.Milk.Price)
	fmt.Println()

	fmt.Printf("Product \"%v\" at price >>$%8.2f<<", models.Milk.Name, models.Milk.Price)
	fmt.Println()

	fmt.Printf("Product \"%v\" at price >>$%.2f<<", models.Milk.Name, models.Milk.Price)
	fmt.Println()
}

func StringsFormatting9() {
	fmt.Println("StringsFormatting9()")

	fmt.Println("StringsFormatting8()")

	fmt.Printf("Product \"%v\" at price >>$%+4.2f<<", models.Milk.Name, models.Milk.Price)
	fmt.Println()

	fmt.Printf("Product \"%v\" at price >>$%08.2f<<", models.Milk.Name, models.Milk.Price)
	fmt.Println()

	fmt.Printf("Product \"%v\" at price >>$%-8.2f<<", models.Milk.Name, models.Milk.Price)
	fmt.Println()
}

func StringsFormatting10() {
	fmt.Println("StringsFormatting10()")

	fmt.Printf("Product \"%s\" at price >>$%4.2f<<", models.Milk.Name, models.Milk.Price)
	fmt.Println()

	fmt.Printf("Product \"%c\" at price >>$%4.2f<<", []rune(models.Milk.Name)[0], models.Milk.Price)
	fmt.Println()

	fmt.Printf("Product \"%U\" at price >>$%4.2f<<", []rune(models.Milk.Name)[0], models.Milk.Price)
	fmt.Println()
}

func StringsFormatting11() {
	fmt.Println("StringsFormatting11()")

	fmt.Printf("Length of milk > 2 = %t", len(models.Milk.Name) > 2)
	fmt.Println()

	fmt.Printf("Length of milk > 100 = %t", len(models.Milk.Name) > 100)
	fmt.Println()
}

func StringsFormatting12() {
	fmt.Println("StringsFormatting12()")

	fmt.Printf("Pointer to milk product %p", models.Milk)
	fmt.Println()
}
