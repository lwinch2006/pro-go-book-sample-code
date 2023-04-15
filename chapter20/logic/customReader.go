package logic

import (
	"chapter20/utils"
	"io"
)

type CustomReader struct {
	reader io.Reader
	count  int
}

func NewCustomReader(reader io.Reader) *CustomReader {
	return &CustomReader{reader, 0}
}

func (cr *CustomReader) Read(slice []byte) (count int, err error) {
	count, err = cr.reader.Read(slice)
	cr.count++
	utils.Printfln("Custom reader has read %v bytes", count)

	if err == io.EOF {
		utils.Printfln("Custom reader total number of reads is %v", cr.count)
	}

	return
}
