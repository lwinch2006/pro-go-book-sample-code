package utils

import (
	"fmt"
	"io"
	"strings"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func ScanFromReader(reader io.Reader, template string, vals ...interface{}) (int, error) {
	return fmt.Fscanf(reader, template, vals...)
}

func ScanSingle(reader io.Reader, vals ...interface{}) (int, error) {
	return fmt.Fscan(reader, vals...)
}

func WriteFormatted(writer io.Writer, template string, vals ...interface{}) (n int, err error) {
	return fmt.Fprintf(writer, template, vals...)
}

func WriteReplaced(writer io.Writer, str string, subs ...string) {
	replacer := strings.NewReplacer(subs...)
	replacer.WriteString(writer, str)
}
