package main

import (
	"chapter25/data"
	"chapter25/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func HttpServer1() {
	utils.Printfln("HttpServer1()")

	http.HandleFunc("/html", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "./httpClients/wwwroot/index.html")
	})

	http.HandleFunc("/json", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(data.Products)
	})

	http.HandleFunc("/echo", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(writer, "Method: %v\n", request.Method)

		for name, values := range request.Header {
			fmt.Fprintf(writer, "Header: %v with value(s) %v\n", name, values)
		}

		fmt.Fprintf(writer, "Body:\n")

		if data, err := io.ReadAll(request.Body); err == nil {
			if len(data) == 0 {
				fmt.Fprintf(writer, "No body in request")
			} else {
				writer.Write(data)
			}
		} else {
			fmt.Fprintf(os.Stdout, "Body reading error: %v", err.Error())
		}
	})

	http.ListenAndServe(":5001", nil)
}

func HttpServer2() {
	utils.Printfln("HttpServer2()")

	http.HandleFunc("/cookie", func(writer http.ResponseWriter, request *http.Request) {
		counter := 1

		if receivedCookie, err := request.Cookie("counter"); err == nil {
			counter, _ = strconv.Atoi(receivedCookie.Value)
			counter++
		}

		http.SetCookie(writer, &http.Cookie{
			Name:  "counter",
			Value: strconv.Itoa(counter),
		})

		if len(request.Cookies()) > 0 {
			for _, c := range request.Cookies() {
				fmt.Fprintf(writer, "Cookie %v with value %v", c.Name, c.Value)
			}
		} else {
			fmt.Fprintf(writer, "Request contains no cookies")
		}
	})

	http.ListenAndServe(":5002", nil)
}

func HttpServer3() {
	utils.Printfln("HttpServer3()")

	http.HandleFunc("/redirect1", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/redirect2", http.StatusTemporaryRedirect)
	})

	http.HandleFunc("/redirect2", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/redirect1", http.StatusTemporaryRedirect)
	})

	http.ListenAndServe(":5003", nil)
}

func HttpServer4() {
	utils.Printfln("HttpServer4()")

	http.HandleFunc("/form", func(writer http.ResponseWriter, request *http.Request) {
		if err := request.ParseMultipartForm(10000000); err == nil {
			for name, values := range request.MultipartForm.Value {
				fmt.Fprintf(writer, "Field %v with values %v\n", name, values)
			}

			for name, files := range request.MultipartForm.File {
				for _, file := range files {
					fmt.Fprintf(writer, "File %v with filename %v\n", name, file.Filename)

					if fileReader, err := file.Open(); err == nil {
						defer fileReader.Close()
						io.Copy(writer, fileReader)
					}
				}
			}
		} else {
			fmt.Fprintf(os.Stdout, "Parsing multipart form error: %v", err.Error())
		}
	})

	http.ListenAndServe(":5004", nil)
}
