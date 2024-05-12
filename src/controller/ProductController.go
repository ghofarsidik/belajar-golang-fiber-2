package controller

import (
	"fmt"
	"gofiber_pijar/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var products = []models.Product{
	{ID: 1, Name: "Cabe", Price: 62000, Stock: 120},
	{ID: 2, Name: "Daging ayam", Price: 25000, Stock: 150},
	{ID: 3, Name: "Gula", Price: 20000, Stock: 200},
}

// menampilkan semua produk
func GetAllProducts(c *fiber.Ctx) error {
	return c.JSON(products)
}

// menampilkan 1 produk
func GetProductById(c *fiber.Ctx) error {
	paramID := c.Params("id")
	id, _ := strconv.Atoi(paramID)

	var foundProduct models.Product
	for _, p := range products {
		if p.ID == id {
			foundProduct = p
			break
		}
	}
	return c.JSON(foundProduct)
}

// menambahkan produk
func CreateProduct(c *fiber.Ctx) error {
	var newProduct models.Product
	if err := c.BodyParser(&newProduct); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
		return err
	}

	newProduct.ID = len(products) + 1

	products = append(products, newProduct)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"product": newProduct,
	})

}

// memperbaharui data produk
func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var updateProduct models.Product
	if err := c.BodyParser(&updateProduct); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Body",
		})
		return err
	}

	var foundIndex int = -1
	for i, p := range products {
		if p.ID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex != -1 {
		products[foundIndex] = updateProduct
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": fmt.Sprintf("Product with ID %d update successfully", id),
			"product": updateProduct,
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Product with ID %d not found", id),
		})
	}
}

// delete product
func DeleteProduct(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	var foundIndex int = -1
	for i, p := range products {
		if p.ID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex != -1 {

		products = append(products[:foundIndex], products[foundIndex+1:]...)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": fmt.Sprintf("Product with ID %d deleted successfully", id),
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Product with ID %d not found", id),
		})
	}
}
