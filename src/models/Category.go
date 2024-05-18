package models

import (
	"gofiber_pijar/src/configs"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string       `json:"name" validate:"required,min=3,max=32"`
	Icon     string       `json:"icon"`
	Products []APIProduct `json:"products"`
}

type APIProduct struct {
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price"`
	Color       string  `json:"color"`
	Size        uint    `json:"size"`
	Stock       uint    `json:"stock"`
	Condition   string  `json:"condition"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	CategoryID  uint    `json:"category_id"`
}

func SelectAllCategory(sort, name string, limit, offset int) []*Category {
	var categories []*Category
	name = "%" + name + "%"
	configs.DB.Preload("Products", func(db *gorm.DB) *gorm.DB {
		var items []*APIProduct
		return db.Model(&Product{}).Find(&items)
	}).Order(sort).Limit(limit).Offset(offset).Where("name ILIKE ?", name).Where("deleted_at IS NULL").Find(&categories)
	return categories
}

func SelectCategoryByID(id int) *Category {
	var item Category
	configs.DB.Preload("Products", func(db *gorm.DB) *gorm.DB {
		var items []*APIProduct
		return db.Model(&Product{}).Find(&items)
	}).First(&item, "id = ?", id)
	return &item
}

func PostCategory(newCategory *Category) error {
	result := configs.DB.Create(&newCategory)
	return result.Error
}

func UpdateCategory(id int, item *Category) error {
	result := configs.DB.Model(&Category{}).Where("id = ?", id).Updates(item)
	return result.Error
}

func DeleteCategory(id int) error {
	result := configs.DB.Delete(&Category{}, "id = ?", id)
	return result.Error
}

func CountDataCategories() int64 {
	var result int64
	configs.DB.Table("categories").Where("deleted_at IS NULL").Count(&result)
	return result
}
