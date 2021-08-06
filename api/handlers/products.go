package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoussa/Ecommerce-Prototype/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	list_of_products := data.GetProducts()
	products_json, err := json.Marshal(list_of_products)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	rw.Write(products_json)
}
