package logic

import (
	"chapter20/utils"
	"io"
)

type CustomWriter struct {
	writer io.Writer
	count  int
}

func NewCustomWriter(writer io.Writer) *CustomWriter {
	return &CustomWriter{writer, 0}
}

func (cw *CustomWriter) Write(slice []byte) (count int, err error) {
	count, err = cw.writer.Write(slice)
	cw.count++
	utils.Printfln("Custom writer has written %v bytes", count)
	return
}

func (cw *CustomWriter) Close() (err error) {
	if closer, ok := cw.writer.(io.Closer); ok {
		closer.Close()
	}

	utils.Printfln("Custom writer total number of writes is %v", cw.count)
	return
}
