package models

import (
	"gofiber_pijar/src/configs"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func SelectAllProduct() []*Product {
	var items []*Product
	configs.DB.Find(&items)
	return items
}

func SelectProductByID(id int) *Product {
	var item Product
	configs.DB.First(&item, "id = ?", id)
	return &item
}

func PostProduct(newProduct *Product) error {
	result := configs.DB.Create(&newProduct)
	return result.Error
}

func UpdateProduct(id int, item *Product) error {
	result := configs.DB.Model(&Product{}).Where("id = ?", id).Updates(item)
	return result.Error
}

func DeleteProduct(id int) error {
	result := configs.DB.Delete(&Product{}, "id = ?", id)
	return result.Error
}
