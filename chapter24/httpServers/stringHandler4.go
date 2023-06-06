package httpServers

import (
	"chapter24/utils"
	"io"
	"net/http"
)

type StringHandler4 struct {
	Message string
}

func (sh StringHandler4) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	utils.Printfln("Requested path: %v", request.URL.Path)

	switch request.URL.Path {
	case "/favicon.ico":
		http.NotFound(writer, request)
	case "/message":
		io.WriteString(writer, sh.Message)
	default:
		http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
	}
}
