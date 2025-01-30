package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	if r.Method == http.MethodPut {
		p.UpdateProduct(w, r)
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
}

func (p *Product) GetProducts(w http.ResponseWriter, r *http.Request) {
	prod := data.GetProducts()

	if err := prod.ToJson(w); err != nil {
		p.l.Println("Error occurred", err)
		return
	}
}

func (p *Product) AddProduct(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	if err := prod.FromJson(r.Body); err != nil {
		p.l.Println("Error unmarshaling product")
		http.Error(w, "Unmarshaling error", http.StatusInternalServerError)
		return
	}

	data.AddProduct(prod)

	p.l.Println("Product added successfully")
}

func (p *Product) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	prod := &data.Product{}

	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error with")
	}

	if err := data.UpdateProduct(i, prod); err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	p.l.Println("Product updated successfully")
}
