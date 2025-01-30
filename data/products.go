package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)

	if err != nil {
		fmt.Println("Error updating products", err)
	}

	p.ID = id
	productList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func (p *Product) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}

func (p *Products) ToJson(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)
}

func AddProduct(p *Product) {
	id := getId()

	p.ID = id

	productList = append(productList, p)
}

func getId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

type Products []*Product

func GetProducts() Products {
	return productList
}

var productList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.UTC.String(),
		UpdatedOn:   time.UTC.String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.UTC.String(),
		UpdatedOn:   time.UTC.String(),
	},
}
