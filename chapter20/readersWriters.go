package main

import (
	"bufio"
	"chapter20/constants"
	"chapter20/logic"
	"chapter20/utils"
	"fmt"
	"io"
	"strings"
)

func ReadersWriters1() {
	fmt.Println("ReadersWriters1()")

	utils.Printfln("Product \"%v\" for price %v", constants.Milk.Name, constants.Milk.Price)
}

func readersWriters2_1(reader io.Reader, writer io.Writer) {
	bytes := make([]byte, 2)

	for {
		count, err := reader.Read(bytes)

		if count > 0 {
			_, _ = writer.Write(bytes[0:count])
			utils.Printfln("%v byte(s) read: %v", count, string(bytes[0:count]))
		}

		if err == io.EOF {
			break
		}
	}

}

func ReadersWriters2() {
	fmt.Println("ReadersWriters2()")

	sr := strings.NewReader(constants.Milk.Name)
	var sw strings.Builder

	readersWriters2_1(sr, &sw)
	utils.Printfln("Written string is %v", sw.String())
}

func readersWriters3_1(reader io.Reader, writer io.Writer) {
	count, err := io.Copy(writer, reader)

	if err == nil {
		utils.Printfln("Number of bytes copied: %v", count)
	} else {
		utils.Printfln("Error occurred: %v", err.Error())
	}

}

func ReadersWriters3() {
	fmt.Println("ReadersWriters3()")

	sr := strings.NewReader(constants.Milk.Name)
	var sw strings.Builder

	readersWriters3_1(sr, &sw)
	utils.Printfln("Written string is %v", sw.String())
}

func readersWriters4_1(writer io.Writer) {
	data := []byte("Milk, Bread")
	writeSize := 4

	for i := 0; i < len(data); i += writeSize {
		end := i + writeSize
		if end > len(data) {
			end = len(data)
		}

		count, err := writer.Write(data[i:end])
		utils.Printfln("Wrote %v byte(s): %v", count, string(data[i:end]))

		if err != nil {
			utils.Printfln("Error occurred: %v", err.Error())
		}
	}

	if closer, ok := writer.(io.Closer); ok {
		_ = closer.Close()
	}
}

func readersWriters4_2(reader io.Reader) {
	data := make([]byte, 0, 10)
	buffer := make([]byte, 2)
	for {
		count, err := reader.Read(buffer)
		if count > 0 {
			utils.Printfln("%v byte(s) read: %v", count, string(buffer[0:count]))
			data = append(data, buffer[0:count]...)
		}

		if err == io.EOF {
			break
		}
	}
}

func ReadersWriters4() {
	fmt.Println("ReadersWriters4()")
	pipeReader, pipeWriter := io.Pipe()

	go readersWriters4_1(pipeWriter)
	readersWriters4_2(pipeReader)
}

func ReadersWriters5() {
	fmt.Println("ReadersWriters5()")
	r1 := strings.NewReader("Milk")
	r2 := strings.NewReader("Bread")
	r3 := strings.NewReader("Beer")

	r4 := io.MultiReader(r1, r2, r3)

	readersWriters4_2(r4)
}

func ReadersWriters6() {
	fmt.Println("ReadersWriters6()")
	var w1 strings.Builder
	var w2 strings.Builder
	var w3 strings.Builder

	w4 := io.MultiWriter(&w1, &w2, &w3)

	readersWriters4_1(w4)

	fmt.Println()
	utils.Printfln("Writer 1: %v", w1.String())
	utils.Printfln("Writer 2: %v", w2.String())
	utils.Printfln("Writer 3: %v", w3.String())
}

func ReadersWriters7() {
	fmt.Println("ReadersWriters7()")
	r1 := strings.NewReader("Milk")
	r2 := strings.NewReader("Bread")
	r3 := strings.NewReader("Beer")

	r4 := io.MultiReader(r1, r2, r3)

	var w1 strings.Builder

	r5 := io.TeeReader(r4, &w1)

	readersWriters4_2(r5)

	utils.Printfln("Writer 1: %v", w1.String())
}

func ReadersWriters8() {
	fmt.Println("ReadersWriters8()")
	r1 := strings.NewReader("Milk")
	r2 := strings.NewReader("Bread")
	r3 := strings.NewReader("Beer")

	r4 := io.MultiReader(r1, r2, r3)
	r5 := io.LimitReader(r4, 5)

	readersWriters4_2(r5)
}

func ReadersWriters9() {
	fmt.Println("ReadersWriters9()")

	text := "It was a boat. A very small boat"

	r1 := logic.NewCustomReader(strings.NewReader(text))
	var w1 strings.Builder

	slice := make([]byte, 5)

	for {
		count, err := r1.Read(slice)
		if count > 0 {
			w1.Write(slice[0:count])
		}

		if err != nil {
			break
		}
	}

	utils.Printfln("Writer data: %v", w1.String())
}

func ReadersWriters10() {
	fmt.Println("ReadersWriters10()")

	text := "It was a boat. A very small boat"

	r1 := logic.NewCustomReader(strings.NewReader(text))
	var w1 strings.Builder

	slice := make([]byte, 5)

	r2 := bufio.NewReader(r1)

	for {
		count, err := r2.Read(slice)
		if count > 0 {
			w1.Write(slice[0:count])
		}

		if err != nil {
			break
		}
	}

	utils.Printfln("Writer data: %v", w1.String())
}

func ReadersWriters11() {
	fmt.Println("ReadersWriters11()")

	text := "It was a boat. A very small boat"

	r1 := logic.NewCustomReader(strings.NewReader(text))
	var w1 strings.Builder

	slice := make([]byte, 5)

	r2 := bufio.NewReader(r1)

	for {
		count, err := r2.Read(slice)
		if count > 0 {
			utils.Printfln("Buffer size: %v, buffered size: %v", r2.Size(), r2.Buffered())
			w1.Write(slice[0:count])
		}

		if err != nil {
			break
		}
	}

	utils.Printfln("Writer data: %v", w1.String())
}

func ReadersWriters12() {
	fmt.Println("ReadersWriters12()")

	text := "It was a boat. A very small boat"

	var w1 strings.Builder
	w2 := logic.NewCustomWriter(&w1)

	for i := 0; true; {
		end := i + 5
		if end > len(text) {
			w2.Write([]byte(text[i:]))
			break
		}

		w2.Write([]byte(text[i:end]))
		i = end
	}

	w2.Close()

	utils.Printfln("Writer data: %v", w1.String())
}

func ReadersWriters13() {
	fmt.Println("ReadersWriters13()")

	text := "It was a boat. A very small boat"

	var w1 strings.Builder
	w2 := logic.NewCustomWriter(&w1)
	w3 := bufio.NewWriterSize(w2, 20)

	for i := 0; true; {
		end := i + 5
		if end > len(text) {
			w3.Write([]byte(text[i:]))
			w3.Flush()
			break
		}

		w3.Write([]byte(text[i:end]))
		i = end
	}

	w2.Close()

	utils.Printfln("Writer data: %v", w1.String())
}

func ReadersWriters14() {
	fmt.Println("ReadersWriters14()")

	text := "Milk Food 23.36 NOK"
	r1 := strings.NewReader(text)
	template := "%s %s %f NOK"

	var name, category string
	var price float64

	_, err := utils.ScanFromReader(r1, template, &name, &category, &price)

	if err != nil {
		utils.Printfln("Error: %v", err.Error())
	} else {
		utils.Printfln("Original text: %v", text)
		utils.Printfln("Product %v with category %v and price %v NOK", name, category, price)
	}
}

func ReadersWriters15() {
	fmt.Println("ReadersWriters15()")

	text := "Milk Food 23.36 NOK"
	r1 := strings.NewReader(text)

	utils.Printfln("Original text: %v", text)

	for {
		var value string
		_, err := utils.ScanSingle(r1, &value)

		if err != nil {
			if err != io.EOF {
				utils.Printfln("Error: %v", err.Error())
			}
			break
		}

		utils.Printfln("Value: %v", value)
	}
}

func ReadersWriters16() {
	fmt.Println("ReadersWriters16()")

	var w1 strings.Builder
	template := "%s %s %.2f NOK"

	utils.WriteFormatted(&w1, template, "Milk", "Food", 23.36)

	utils.Printfln("Writer: %v", w1.String())
}

func ReadersWriters17() {
	fmt.Println("ReadersWriters17()")

	text := "It was a boat. A very big boat"
	var w1 strings.Builder

	utils.WriteReplaced(&w1, text, "boat", "yacht", "big", "small")

	utils.Printfln("Original text: %v", text)
	utils.Printfln("Replaced text: %v", w1.String())
}
