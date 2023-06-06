package models

import "net/http"

type ProductListViewModel struct {
	Data    []*Product
	Request *http.Request
}
