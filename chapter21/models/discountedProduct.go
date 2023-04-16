package models

import (
	"encoding/json"
	"strconv"
)

type DiscountedProduct struct {
	//*Product `json:"Product,omitempty"` // This will include Product as object
	*Product `json:",omitempty"` // This will omit Product as object and include properties as flat list/map
	//Discount float64 `json:"-"` // Discount field will be omitted when serializing to JSON
	Discount float64 `json:"Offer,string"` // Discount field will be serialized as string to JSON
}

func NewDiscountedProduct(name, category string, price, discount float64) *DiscountedProduct {
	return &DiscountedProduct{NewProduct(name, category, price), discount}
}

func (p *DiscountedProduct) GetName() string {
	return p.Name
}

// Change name from MarshalJSON123() to MarshalJSON() to see custom serialization to JSON
func (p *DiscountedProduct) MarshalJSON123() (jsn []byte, err error) {
	if p != nil {
		m := map[string]interface{}{
			"product": p.Name,
			"cost":    p.Price - p.Discount,
		}

		jsn, err = json.Marshal(m)
	}

	return
}

func (p *DiscountedProduct) UnmarshalJSON(data []byte) (err error) {
	mdata := map[string]interface{}{}
	err = json.Unmarshal(data, &mdata)

	if p.Product == nil {
		p.Product = &Product{}
	}

	if err != nil {
		return
	}

	if name, ok := mdata["Name"].(string); ok {
		p.Product.Name = name
	}

	if category, ok := mdata["Category"].(string); ok {
		p.Product.Category = category
	}

	if price, ok := mdata["Price"].(float64); ok {
		p.Product.Price = price
	}

	if discountAsString, ok := mdata["Offer"].(string); ok {
		if discount, parseErr := strconv.ParseFloat(discountAsString, 64); parseErr == nil {
			p.Discount = discount
		}
	}

	return
}
