package main

import (
	"bytes"
	"chapter25/data"
	"chapter25/models"
	"chapter25/utils"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

func HttpClient1() {
	utils.Printfln("HttpClient1()")

	if response, err := http.Get("http://localhost:5001/html"); err == nil {
		response.Write(os.Stdout)
	} else {
		utils.Printfln("Sending GET request error: %v", err.Error())
	}
}

func HttpClient2() {
	utils.Printfln("HttpClient2()")

	if response, err := http.Get("http://localhost:5001/html"); err == nil {
		if response.StatusCode == http.StatusOK {
			defer response.Body.Close()
			if data, err := io.ReadAll(response.Body); err == nil {
				os.Stdout.Write(data)
			}
		} else {
			utils.Printfln("Error HTTP status code: %v", response.StatusCode)
		}
	} else {
		utils.Printfln("Sending GET request error: %v", err.Error())
	}
}

func HttpClient3() {
	utils.Printfln("HttpClient3()")

	if response, err := http.Get("http://localhost:5001/html123"); err == nil {
		if response.StatusCode == http.StatusOK {
			defer response.Body.Close()
			if data, err := io.ReadAll(response.Body); err == nil {
				os.Stdout.Write(data)
			}
		} else {
			utils.Printfln("Error HTTP status code: %v", response.StatusCode)
		}
	} else {
		utils.Printfln("Sending GET request error: %v", err.Error())
	}
}

func HttpClient4() {
	utils.Printfln("HttpClient4()")

	if response, err := http.Get("http://localhost:5001/json"); err == nil {
		if response.StatusCode == http.StatusOK {
			defer response.Body.Close()

			var products []*models.Product
			json.NewDecoder(response.Body).Decode(&products)

			for _, p := range products {
				utils.Printfln("Product: %v, category: %v, price: %.2f", p.Name, p.Category, p.Price)
			}
		} else {
			utils.Printfln("Error HTTP status code: %v", response.StatusCode)
		}
	} else {
		utils.Printfln("Sending GET request error: %v", err.Error())
	}
}

func HttpClient5() {
	utils.Printfln("HttpClient5()")

	formData := map[string][]string{
		"name":     {"Vodka"},
		"category": {"Food"},
		"price":    {"101.23"},
	}

	if response, err := http.PostForm("http://localhost:5001/echo", formData); err == nil {
		if response.StatusCode == http.StatusOK {
			defer response.Body.Close()
			io.Copy(os.Stdout, response.Body)
		} else {
			utils.Printfln("Error HTTP status code: %v", response.StatusCode)
		}
	} else {
		utils.Printfln("Sending POST request error: %v", err.Error())
	}
}

func HttpClient6() {
	utils.Printfln("HttpClient6()")

	var sb strings.Builder
	json.NewEncoder(&sb).Encode(data.Products)

	if response, err := http.Post("http://localhost:5001/echo", "application/json", strings.NewReader(sb.String())); err == nil {
		if response.StatusCode == http.StatusOK {
			defer response.Body.Close()
			io.Copy(os.Stdout, response.Body)
		} else {
			utils.Printfln("Error HTTP status code: %v", response.StatusCode)
		}
	} else {
		utils.Printfln("Sending POST request error: %v", err.Error())
	}
}

func HttpClient7() {
	utils.Printfln("HttpClient7()")

	var sb strings.Builder
	json.NewEncoder(&sb).Encode(data.Products)

	requestUrl, _ := url.Parse("http://localhost:5001/echo")

	request := &http.Request{
		Method: "POST",
		URL:    requestUrl,
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
		Body: io.NopCloser(strings.NewReader(sb.String())),
	}

	if response, err := http.DefaultClient.Do(request); err == nil {
		if response.StatusCode == http.StatusOK {
			defer response.Body.Close()
			io.Copy(os.Stdout, response.Body)
		} else {
			utils.Printfln("Error HTTP status code: %v", response.StatusCode)
		}
	} else {
		utils.Printfln("Sending POST request error: %v", err.Error())
	}
}

func HttpClient8() {
	utils.Printfln("HttpClient8()")

	var sb strings.Builder
	json.NewEncoder(&sb).Encode(data.Products)

	// without io.NopCloser also works and Content-Length header will be added to the request
	request, _ := http.NewRequest("POST", "http://localhost:5001/echo", io.NopCloser(strings.NewReader(sb.String())))
	request.Header.Set("Content-Type", "application/json")

	if response, err := http.DefaultClient.Do(request); err == nil {
		if response.StatusCode == http.StatusOK {
			defer response.Body.Close()
			io.Copy(os.Stdout, response.Body)
		} else {
			utils.Printfln("Error HTTP status code: %v", response.StatusCode)
		}
	} else {
		utils.Printfln("Sending POST request error: %v", err.Error())
	}
}

func HttpClient9() {
	utils.Printfln("HttpClient9()")

	cookieJar, _ := cookiejar.New(nil)
	http.DefaultClient.Jar = cookieJar

	for i := 0; i < 3; i++ {
		request, _ := http.NewRequest(http.MethodGet, "http://localhost:5002/cookie", nil)
		if response, err := http.DefaultClient.Do(request); err == nil {
			if response.StatusCode == http.StatusOK {
				defer response.Body.Close()
				io.Copy(os.Stdout, response.Body)
				fmt.Println()
			} else {
				utils.Printfln("Error HTTP status code: %v", response.StatusCode)
			}
		} else {
			utils.Printfln("Sending POST request error: %v", err.Error())
		}
	}
}

func HttpClient10() {
	utils.Printfln("HttpClient10()")

	clients := make([]http.Client, 3)

	for index, client := range clients {

		// Separate cookies for each http client
		cookieJar, _ := cookiejar.New(nil)
		client.Jar = cookieJar

		for i := 0; i < 3; i++ {
			request, _ := http.NewRequest(http.MethodGet, "http://localhost:5002/cookie", nil)
			if response, err := client.Do(request); err == nil {
				if response.StatusCode == http.StatusOK {
					defer response.Body.Close()
					fmt.Print("Client ", index, " ")
					io.Copy(os.Stdout, response.Body)
					fmt.Println()
				} else {
					utils.Printfln("Error HTTP status code: %v", response.StatusCode)
				}
			} else {
				utils.Printfln("Sending POST request error: %v", err.Error())
			}
		}
	}
}

func HttpClient11() {
	utils.Printfln("HttpClient11()")

	clients := make([]http.Client, 3)
	cookieJar, _ := cookiejar.New(nil)

	for index, client := range clients {

		// shared cookies for all http clients
		client.Jar = cookieJar

		for i := 0; i < 3; i++ {
			request, _ := http.NewRequest(http.MethodGet, "http://localhost:5002/cookie", nil)
			if response, err := client.Do(request); err == nil {
				if response.StatusCode == http.StatusOK {
					defer response.Body.Close()
					fmt.Print("Client ", index, " ")
					io.Copy(os.Stdout, response.Body)
					fmt.Println()
				} else {
					utils.Printfln("Error HTTP status code: %v", response.StatusCode)
				}
			} else {
				utils.Printfln("Sending POST request error: %v", err.Error())
			}
		}
	}
}

func HttpClient12() {
	utils.Printfln("HttpClient12()")

	http.DefaultClient.CheckRedirect = func(request *http.Request, requests []*http.Request) error {
		if len(requests) >= 3 {
			request.URL, _ = url.Parse("http://localhost:5001/html")
		}

		return nil
	}

	if response, err := http.Get("http://localhost:5003/redirect1"); err == nil {
		if response.StatusCode == http.StatusOK {
			defer response.Body.Close()
			if data, err := io.ReadAll(response.Body); err == nil {
				os.Stdout.Write(data)
			}
		} else {
			utils.Printfln("Error HTTP status code: %v", response.StatusCode)
		}
	} else {
		utils.Printfln("Sending GET request error: %v", err.Error())
	}
}

func HttpClient13() {
	utils.Printfln("HttpClient13()")

	var buffer bytes.Buffer
	formWriter := multipart.NewWriter(&buffer)
	fieldWriter, _ := formWriter.CreateFormField("name")
	io.WriteString(fieldWriter, "Alice")
	fieldWriter, _ = formWriter.CreateFormField("city")
	io.WriteString(fieldWriter, "New York")
	fieldWriter, _ = formWriter.CreateFormFile("file", "customFmt.go")
	fileData, _ := os.ReadFile("./utils/customFmt.go")
	fieldWriter.Write(fileData)
	formWriter.Close()

	if response, err := http.Post("http://localhost:5004/form", formWriter.FormDataContentType(), &buffer); err == nil {
		if response.StatusCode == http.StatusOK {
			defer response.Body.Close()
			if data, err := io.ReadAll(response.Body); err == nil {
				os.Stdout.Write(data)
			}
		} else {
			utils.Printfln("Error HTTP status code: %v", response.StatusCode)
		}
	} else {
		utils.Printfln("Sending GET request error: %v", err.Error())
	}
}
