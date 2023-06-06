package httpServers

import (
	"chapter24/data"
	"chapter24/models"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

var HtmlTemplates *template.Template

func HandleHtmlTemplate(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path

	if path == "" {
		path = "products.html"
	}

	t := HtmlTemplates.Lookup(path)
	if t == nil {
		http.NotFound(writer, request)
	} else {
		viewModel := models.ProductListViewModel{
			Data:    data.Products,
			Request: request,
		}

		//HtmlTemplates.ExecuteTemplate(writer, path, viewModel) // Alternative way to execute template

		err := t.Execute(writer, viewModel)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}

func ProcessFormData(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		http.Redirect(writer, request, "/templates", http.StatusTemporaryRedirect)
		return
	}

	productIndex, _ := strconv.Atoi(request.PostFormValue("index"))
	updatedProduct := &models.Product{}
	updatedProduct.Name = request.PostFormValue("name")
	updatedProduct.Category = request.PostFormValue("category")
	updatedProduct.Price, _ = strconv.ParseFloat(request.PostFormValue("price"), 64)
	data.Products[productIndex] = updatedProduct

	http.Redirect(writer, request, "/templates", http.StatusTemporaryRedirect)
}

func UploadSingleFileHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		http.Redirect(writer, request, "/templates", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(writer, "Name: %v, City: %v\n", request.PostFormValue("name"), request.PostFormValue("city"))
	fmt.Fprintln(writer, "-------")

	file, header, err := request.FormFile("files")

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()
	fmt.Fprintf(writer, "Name: %v, Size: %v\n", header.Filename, header.Size)

	for k, v := range header.Header {
		fmt.Fprintf(writer, "Key: %v, Value: %v\n", k, v)
	}

	fmt.Fprintln(writer, "------")
	io.Copy(writer, file)
}

func UploadMultipleFileHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		http.Redirect(writer, request, "/templates", http.StatusTemporaryRedirect)
		return
	}

	request.ParseMultipartForm(10000000)

	fmt.Fprintf(writer, "Name: %v, City: %v\n",
		request.MultipartForm.Value["name"][0],
		request.MultipartForm.Value["city"][0])
	fmt.Fprintln(writer, "--------")

	for _, header := range request.MultipartForm.File["files"] {
		fmt.Fprintf(writer, "File name: %v, File size: %v\n", header.Filename, header.Size)

		for k, v := range header.Header {
			fmt.Fprintf(writer, "Key: %v, Value: %v\n", k, v)
		}

		file, err := header.Open()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()
		fmt.Fprintln(writer, "File content:")
		io.Copy(writer, file)
		fmt.Fprintln(writer, "")
		fmt.Fprintln(writer, "--------")
	}
}
