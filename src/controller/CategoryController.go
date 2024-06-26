package controller

import (
	"fmt"
	"gofiber_pijar/src/helpers"
	"gofiber_pijar/src/models"
	"math"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
)

// menampilkan semua kategori
func GetAllCategories(c *fiber.Ctx) error {

	//pagination
	pageParam := c.Query("page")
	limitParam := c.Query("limit")

	page, _ := strconv.Atoi(pageParam)
	if page == 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(limitParam)
	if limit == 0 {
		limit = 5
	}

	offset := (page - 1) * limit

	//sorting
	sort := c.Query("sorting")
	if sort == "" {
		sort = "ASC"
	}

	sortBy := c.Query("sortBy")
	if sortBy == "" {
		sortBy = "name"
	}

	sort = sortBy + " " + strings.ToLower(sort)

	keyword := c.Query("search")
	categories := models.SelectAllCategory(sort, keyword, limit, offset)
	totalData := models.CountDataCategories()
	totalPage := math.Ceil(float64(totalData) / float64(limit))
	result := map[string]interface{}{
		"data":        categories,
		"currentPage": page,
		"limit":       limit,
		"totalData":   totalData,
		"totalPage":   totalPage,
	}

	return c.JSON(result)
}

// menampilkan 1 kategori sesuai ID
func GetCategoryById(c *fiber.Ctx) error {
	paramID := c.Params("id")
	id, _ := strconv.Atoi(paramID)

	foundCategory := models.SelectCategoryByID(id)

	if foundCategory == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Category not found",
		})
	}
	return c.JSON(foundCategory)
}

// menambah kategori
func CreateCategory(c *fiber.Ctx) error {
	var rawNewCategory map[string]interface{}
	if err := c.BodyParser(&rawNewCategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// xssmiddleware
	rawNewCategory = helpers.XSSMiddleware(rawNewCategory)

	var newCategory models.Category
	mapstructure.Decode(rawNewCategory, &newCategory)

	//validate
	errors := helpers.ValidateStruct(newCategory)
	if len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	models.PostCategory(&newCategory)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Category created successfully",
	})
}

// memperbaharui kategori
func UpdateCategory(c *fiber.Ctx) error {
	paramID := c.Params("id")
	id, _ := strconv.Atoi(paramID)

	var rawUpdateCategory map[string]interface{}

	if err := c.BodyParser(&rawUpdateCategory); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	//XSSMiddleware
	rawUpdateCategory = helpers.XSSMiddleware(rawUpdateCategory)

	var UpdateCategory models.Category
	mapstructure.Decode(rawUpdateCategory, &UpdateCategory)

	//validate
	errors := helpers.ValidateStruct(UpdateCategory)
	if len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	err := models.UpdateCategory(id, &UpdateCategory)

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": fmt.Sprintf("Category with ID %d updated successfully", id),
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Category with ID %d not found", id),
		})
	}
}

// delete category
func DeleteCategory(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	err := models.DeleteCategory(id)

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": fmt.Sprintf("Category with ID %d deleted successfully", id),
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Product with ID %d not found", id),
		})
	}
}
