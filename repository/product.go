package repository

import (
	"go-crud-gorm/database"
	"go-crud-gorm/entity"
	"go-crud-gorm/logger"
)

func CreateProduct(product entity.Product) {
	logger.Info.Print("Starting create product")

	database.Db.Create(&product)

	logger.Info.Print("Product created")
}

func CreateProducts(products []entity.Product) {
	logger.Info.Print("Starting create products")

	database.Db.Create(&products)

	logger.Info.Print("Products created")
}

func SelectProductByPk(id string) entity.Product {
	logger.Info.Print("Starting select product")

	var product entity.Product
	database.Db.First(&product, "id = ?", id)

	logger.Info.Print("Product selected!")

	return product
}

func SelectAllProducts() []entity.Product {
	logger.Info.Print("Starting select all products")

	var products []entity.Product
	database.Db.Find(&products)

	logger.Info.Print("All Products selecteds")

	return products
}

func UpdateProduct(id string, product entity.Product) {
	logger.Info.Print("Starting update product")

	var selectedProduct entity.Product
	database.Db.First(&selectedProduct, "id = ?", id)

	selectedProduct = product
	database.Db.Save(&selectedProduct)

	logger.Info.Print("Product updated")
}

func DeleteProduct(id string) {
	logger.Info.Print("Starting delete product")

	var selectedProduct entity.Product
	database.Db.First(&selectedProduct, "id = ?", id)

	database.Db.Delete(&selectedProduct)

	logger.Info.Print("Product deleted")
}
