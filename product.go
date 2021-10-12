package eplaza

import (
	"github.com/google/uuid"
)

type Product struct {
	Id                      uuid.UUID
	Name, Price, Product_Id string

	Photos                 []string
	Created_at, Updated_at uuid.Time
}

type ProductService interface {
	//create Product
	CreateProduct() error
	//get Product
	GetProduct(id int) Product
	//Get all Products
	GetAllProducts() []Product
	//update Product
	UpdateProduct(id int) error
	//delete Product
	DeleteProduct(id int) error
}
