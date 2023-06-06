package httpServers

import (
	"chapter24/utils"
	"io"
	"net/http"
)

type StringHandler struct {
	Message string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	utils.Printfln("Method: %v", request.Method)
	utils.Printfln("URL: %v", request.URL)
	utils.Printfln("Proto: %v", request.Proto)
	utils.Printfln("Host: %v", request.Host)

	for name, values := range request.Header {
		utils.Printfln("Header name: %v, header values: %v", name, values)
	}

	utils.Printfln("---")

	io.WriteString(writer, sh.Message)
}
