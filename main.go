package main

import (
	"go-crud-gorm/database"
	"go-crud-gorm/entity"
	"go-crud-gorm/logger"
	"go-crud-gorm/repository"
)

func init() {
	err := logger.MakeLoggers()
	if err != nil {
		panic(err)
	}

	err = database.Connect()
	if err != nil {
		panic(err)
	}
}

func main() {
	logger.Info.Print("Application started")

	product1 := entity.NewProduct("first product", 100)
	repository.CreateProduct(*product1)

	product2 := entity.NewProduct("second product", 200)
	product3 := entity.NewProduct("third product", 300)

	var products []entity.Product
	products = append(products, *product2, *product3)
	repository.CreateProducts(products)

	selectedProduct := repository.SelectProductByPk(product1.ID)
	logger.Info.Printf("Product1: %v", selectedProduct)

	product1.Name = "First product updated"
	repository.UpdateProduct(product1.ID, *product1)

	repository.DeleteProduct(product2.ID)

	selectedProducts := repository.SelectAllProducts()
	for _, selectedProduct := range selectedProducts {
		logger.Info.Printf("All Products: %v", selectedProduct)
	}

	logger.Info.Print("Application finished")
}
