package httpServers

import (
	"chapter24/data"
	"encoding/json"
	"net/http"
)

func JsonHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(data.Products)
}
