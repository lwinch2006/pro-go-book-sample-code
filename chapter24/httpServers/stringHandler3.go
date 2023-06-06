package httpServers

import (
	"chapter24/utils"
	"io"
	"net/http"
)

type StringHandler3 struct {
	Message string
}

func (sh StringHandler3) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/favicon.ico" {
		utils.Printfln("Requesting favicon - return 404")
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	utils.Printfln("Requested path: %v", request.URL.Path)
	io.WriteString(writer, sh.Message)
}
