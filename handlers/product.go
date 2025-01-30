package handlers

import (
	"log"
	"net/http"

	"main.go/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.AddProduct(w, r)
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
}

func (p *Product) GetProducts(w http.ResponseWriter, r *http.Request) {
	prod := data.GetProducts()

	if err := prod.FromJson(w); err != nil {
		p.l.Println("Error occurred", err)
		return
	}
}

func (p *Product) AddProduct(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	if err := prod.ToJson(r.Body); err != nil {
		p.l.Println("Error unmarshaling product")
		http.Error(w, "Unmarshaling error", http.StatusInternalServerError)
		return
	}

	data.AddProduct(prod)

	p.l.Println("Product added successfully")
}
