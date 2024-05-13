package controller

import (
	"fmt"
	"gofiber_pijar/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// menampilkan semua produk
func GetAllProducts(c *fiber.Ctx) error {
	products := models.SelectAllProduct()
	return c.JSON(products)
}

// menampilkan 1 produk
func GetProductById(c *fiber.Ctx) error {
	paramID := c.Params("id")
	id, _ := strconv.Atoi(paramID)

	foundProduct := models.SelectProductByID(id)

	if foundProduct == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	return c.JSON(foundProduct)
}

// // menambahkan produk
func CreateProduct(c *fiber.Ctx) error {
	var newProduct models.Product
	if err := c.BodyParser(&newProduct); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
		return err
	}

	models.PostProduct(&newProduct)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
	})

}

// // memperbaharui data produk
func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var updateProduct models.Product
	if err := c.BodyParser(&updateProduct); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Body",
		})
		return err
	}

	err := models.UpdateProduct(id, &updateProduct)

	if err == nil {
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

// // delete product
func DeleteProduct(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	err := models.DeleteProduct(id)

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": fmt.Sprintf("Product with ID %d deleted successfully", id),
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Product with ID %d not found", id),
		})
	}
}
