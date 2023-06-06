package httpServers

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetAndSetCookies(writer http.ResponseWriter, request *http.Request) {
	counter := 1
	counterCookie, err := request.Cookie("counter")
	if err == nil {
		counter, _ = strconv.Atoi(counterCookie.Value)
		counter++
	}

	http.SetCookie(writer, &http.Cookie{
		Name:  "counter",
		Value: strconv.Itoa(counter),
	})

	if len(request.Cookies()) > 0 {
		for _, cookie := range request.Cookies() {
			fmt.Fprintf(writer, "Cookie name: %v, Cookie value: %v\n", cookie.Name, cookie.Value)
		}
	} else {
		fmt.Fprintln(writer, "Request contains no cookies")
	}
}
