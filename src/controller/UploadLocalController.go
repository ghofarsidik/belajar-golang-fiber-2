package controller

import (
	"fmt"
	"gofiber_pijar/src/helpers"

	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Gagal mengunggah file: " + err.Error())
	}

	maxFileSize := int64(2 << 20)
	if err := helpers.SizeUploadValidation(file.Size, maxFileSize); err != nil {
		return err
	}

	fileHeader, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal membuka file: " + err.Error())
	}
	defer fileHeader.Close()

	buffer := make([]byte, 512)
	if _, err := fileHeader.Read(buffer); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal membaca file: " + err.Error())
	}

	validFileTypes := []string{"image/png", "image/jpeg", "image/jpg", "application/pdf"}
	if err := helpers.TypeUploadValidation(buffer, validFileTypes); err != nil {
		return err
	}

	filePath := helpers.UploadFIle(file)
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal menyimpan file " + err.Error())
	}

	return c.SendString(fmt.Sprintf("File %s berhasil di upload ke %s", file.Filename, filePath))
}
