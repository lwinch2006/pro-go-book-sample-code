package httpServers

import (
	"chapter24/utils"
	"io"
	"net/http"
	"strings"
)

type StringHandler7 struct {
	Message string
}

func (sh StringHandler7) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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

func RedirectToHttps(writer http.ResponseWriter, request *http.Request) {
	host := strings.Split(request.Host, ":")[0]
	target := "https://" + host + ":5557" + request.URL.Path

	if len(request.URL.RawQuery) > 0 {
		target += "?" + request.URL.RawQuery
	}

	http.Redirect(writer, request, target, http.StatusTemporaryRedirect)
}
