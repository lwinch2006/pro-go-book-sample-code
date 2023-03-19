package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func Strings1() {
	fmt.Println("Strings1()")

	product := "Milk"

	// These comparisons are case-sensitive, except EqualFold()

	fmt.Println("Product:", product)

	fmt.Println("Contains(Milk, lk):", strings.Contains(product, "lk"))
	fmt.Println("ContainsAny(Milk, xyz):", strings.ContainsAny(product, "xyz"))
	fmt.Println("ContainsRune(Milk, M):", strings.ContainsRune(product, 'M'))
	fmt.Println("EqualFold(Milk, mIlK):", strings.EqualFold(product, "mIlK"))
	fmt.Println("HasPrefix(Milk, Br):", strings.HasPrefix(product, "Br"))
	fmt.Println("HasSuffix(Milk, ilk):", strings.HasSuffix(product, "ilk"))

	fmt.Println("HasSuffix(Milk, ilk):", bytes.HasSuffix([]byte(product), []byte{105, 108, 107}))
}

func Strings2() {
	fmt.Println("Strings2()")

	description := "Milk is for drinking"

	fmt.Println("Title(str):", strings.Title(description))
	fmt.Println("ToTitle(str):", strings.ToTitle(description))
}

func Strings3() {
	fmt.Println("Strings3()")

	specialUnicode := "\u01c9"
	upperSpecialUnicode := strings.ToUpper(specialUnicode)
	toTitleSpecialUnicode := strings.ToTitle(specialUnicode)

	fmt.Println("Original:", specialUnicode, []byte(specialUnicode))
	fmt.Println("Upper", upperSpecialUnicode, []byte(upperSpecialUnicode))
	fmt.Println("ToTitle:", toTitleSpecialUnicode, []byte(toTitleSpecialUnicode))
}

func Strings4() {
	fmt.Println("Strings4()")

	product := "Milk"

	for _, chr := range product {
		fmt.Println("Char:", string(chr), "IsLower():", unicode.IsLower(chr), "ToLower():", string(unicode.ToLower(chr)), "IsUpper():", unicode.IsUpper(chr), "ToUpper():", string(unicode.ToUpper(chr)), "IsTitle():", unicode.IsTitle(chr), "ToTitle():", string(unicode.ToTitle(chr)))
	}
}

func Strings5() {
	fmt.Println("Strings5()")

	description := "Milk is for drinking"

	// These comparisons are case-sensitive

	fmt.Println("Original:", description)
	fmt.Println("Count(in):", strings.Count(description, "in"))
	fmt.Println("Index(in):", strings.Index(description, "in"))
	fmt.Println("LastIndex(in):", strings.LastIndex(description, "in"))

	fmt.Println("IndexAny(in):", strings.IndexAny(description, "in"))
	fmt.Println("LastIndexAny(in):", strings.LastIndexAny(description, "in"))

	fmt.Println("IndexByte(i):", strings.IndexByte(description, 'i'))
	fmt.Println("LastIndexByte(i):", strings.LastIndexByte(description, 'i'))
}

func strings6_1(source rune) bool {
	return source == 'k' || source == 'K'
}

func Strings6() {
	fmt.Println("Strings6()")

	description := "Milk is for drinking"

	fmt.Println("Original:", description)
	fmt.Println("IndexFunc(k):", strings.IndexFunc(description, strings6_1))
	fmt.Println("LastIndexFunc(k):", strings.LastIndexFunc(description, strings6_1))
}

func Strings7() {
	fmt.Println("Strings7()")

	description := "Milk is for drinking"
	split1 := strings.Fields(description)
	split2 := strings.Split(description, " ")
	split3 := strings.SplitAfter(description, " ")

	fmt.Println("Split by fields")
	for _, item := range split1 {
		fmt.Println(">>" + item + "<<")
	}

	fmt.Println("Split by whitespaces")
	for _, item := range split2 {
		fmt.Println(">>" + item + "<<")
	}

	fmt.Println("Split after by whitespaces")
	for _, item := range split3 {
		fmt.Println(">>" + item + "<<")
	}
}

func Strings8() {
	fmt.Println("Strings8()")

	description := "Milk is for drinking"
	split1 := strings.SplitN(description, " ", 3)
	split2 := strings.SplitAfterN(description, " ", 3)

	fmt.Println("SplitN by whitespaces")
	for _, item := range split1 {
		fmt.Println(">>" + item + "<<")
	}

	fmt.Println("Split after N by whitespaces")
	for _, item := range split2 {
		fmt.Println(">>" + item + "<<")
	}
}

func Strings9() {
	fmt.Println("Strings9()")

	description := "Milk  is  for  drinking"
	split1 := strings.Fields(description)
	split2 := strings.Split(description, " ")
	split3 := strings.SplitAfter(description, " ")

	fmt.Println("Split (double spaces) by fields")
	for _, item := range split1 {
		fmt.Println(">>" + item + "<<")
	}

	fmt.Println("Split (double spaces) by whitespaces")
	for _, item := range split2 {
		fmt.Println(">>" + item + "<<")
	}

	fmt.Println("Split after (double spaces) by whitespaces")
	for _, item := range split3 {
		fmt.Println(">>" + item + "<<")
	}
}

func strings10_1(chr rune) bool {
	return chr == ' '
}

func Strings10() {
	fmt.Println("Strings10()")

	description := "Milk  is  for  drinking"
	split1 := strings.FieldsFunc(description, strings10_1)

	fmt.Println("Split (double spaces) by fields func")
	for _, item := range split1 {
		fmt.Println(">>" + item + "<<")
	}
}

func strings11_1(chr rune) bool {
	return chr == ' ' || chr == 'M' || chr == 'g'
}

func Strings11() {
	fmt.Println("Strings11()")

	description := " Milk  is  for  drinking "

	fmt.Println("Original:", ">>"+description+"<<")
	fmt.Println("TrimSpace():", ">>"+strings.TrimSpace(description)+"<<")
	fmt.Println("Trim(\" \"):", ">>"+strings.Trim(description, " ")+"<<")
	fmt.Println("TrimLeft(\" \"):", ">>"+strings.TrimLeft(description, " ")+"<<")
	fmt.Println("TrimRight(\" \"):", ">>"+strings.TrimRight(description, " ")+"<<")
	fmt.Println("Trim(\" Mg\"):", ">>"+strings.Trim(description, " Mg")+"<<")

	fmt.Println("TrimPrefix(\" Milk\"):", ">>"+strings.TrimPrefix(description, " Milk")+"<<")
	fmt.Println("TrimSuffix(\"ing \"):", ">>"+strings.TrimSuffix(description, "ing ")+"<<")

	fmt.Println("TrimFunc():", ">>"+strings.TrimFunc(description, strings11_1)+"<<")
	fmt.Println("TrimLeftFunc():", ">>"+strings.TrimLeftFunc(description, strings11_1)+"<<")
	fmt.Println("TrimRightFunc():", ">>"+strings.TrimRightFunc(description, strings11_1)+"<<")

}

func strings12_1(chr rune) rune {
	if chr == ' ' {
		return '*'
	}

	return chr
}

func Strings12() {
	fmt.Println("Strings12()")

	description := " Milk  is  for  drinking "

	fmt.Println("Original:", ">>"+description+"<<")
	fmt.Println("Replace(\" \", 2):", strings.Replace(description, " ", "*", 2))
	fmt.Println("ReplaceAll():", strings.ReplaceAll(description, " ", "*"))
	fmt.Println("Map():", strings.Map(strings12_1, description))

	replacer := strings.NewReplacer("Milk", "Bread", "drinking", "eating")

	fmt.Println("Replace with replacer:", replacer.Replace(description))
}

func Strings13() {
	fmt.Println("Strings13()")

	description := "Milk  is  for  drinking"
	slice1 := strings.Fields(description)

	fmt.Println("Original:", ">>"+description+"<<")
	fmt.Println("Join(\"--\"):", strings.Join(slice1, "--"))
	fmt.Println("Repeat():", strings.Repeat(slice1[0]+" ", 3))
}

func Strings14() {
	fmt.Println("Strings14()")

	description := "Milk is for drinking."
	slice1 := strings.Fields(description)

	var strBuilder strings.Builder

	for _, item := range slice1 {
		if item == "for" {
			strBuilder.WriteString("to")
		} else if item == "drinking." {
			strBuilder.WriteString("drink.")
		} else {
			strBuilder.WriteString(item)
		}

		strBuilder.WriteRune(' ')
	}

	fmt.Println("Original:", ">>"+description+"<<")
	fmt.Println("Built string:", ">>"+strBuilder.String()+"<<")
}
