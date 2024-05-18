package models

import (
	"gofiber_pijar/src/configs"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string   `json:"name" validate:"required"`
	Image       string   `json:"image" validate:"required"`
	Brand       string   `json:"brand" validate:"required"`
	Price       float64  `json:"price" validate:"required"`
	Color       string   `json:"color" validate:"required"`
	Size        uint     `json:"size" validate:"required"`
	Stock       uint     `json:"stock" validate:"required"`
	Condition   string   `json:"condition" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Rating      float32  `json:"rating"`
	CategoryId  uint     `json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryId"`
}

//preload buat ambil semua, select buat ambil beberapa

func SelectAllProduct(sort, name string, limit, offset int) []*Product {
	var items []*Product
	name = "%" + name + "%"
	configs.DB.Preload("Category").Order(sort).Limit(limit).Offset(offset).Where("name ILIKE ?", name).Where("deleted_at IS NULL").Find(&items)
	return items
}

func SelectProductByID(id int) *Product {
	var item Product
	configs.DB.Preload("Category").First(&item, "id = ?", id)
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

func CountDataProducts() int64 {
	var result int64
	configs.DB.Table("products").Where("deleted_at IS NULL").Count(&result)
	return result
}
