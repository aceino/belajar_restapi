package controllers

import (
	"strconv"

	"github.com/aceino/belajar_restapi/config"
	"github.com/aceino/belajar_restapi/models"
	"github.com/gofiber/fiber/v2"
)

func InsertBook(c *fiber.Ctx) error {
	db := config.DB

	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.JSON(fiber.Map{"Error": err})
	}

	db.Create(&book)

	return c.JSON(fiber.Map{"data": book})

}

func GetAllBook(c *fiber.Ctx) error {
	db := config.DB

	var book []models.Book

	db.Find(&book)

	if len(book) == 0 {
		return c.Status(404).JSON(fiber.Map{"data": book})
	}

	return c.Status(200).JSON(fiber.Map{"data": book})

}

func GetBookById(c *fiber.Ctx) error {

	db := config.DB

	var book []models.Book

	// book := new(models.Book)

	param := c.Params("id")

	//  konversi tipe data
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Data is not found",
		})
	}

	db.First(&book, id)

	return c.Status(200).JSON(fiber.Map{"data": book})

}

func UpdateBook(c *fiber.Ctx) error {

	db := config.DB

	book := new(models.Book)
	// var book []models.Book

	params := c.Params("id")

	// konversi tipe data
	id, err := strconv.Atoi(params)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Data is not fond",
		})
	}

	// Binds the request body to a struct.
	if err := c.BodyParser(book); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	db.Where("id=?", id).Updates(&book)
	
	return c.Status(200).JSON(fiber.Map{"data": book})
}

func DeleteBook(c *fiber.Ctx) error {

	db := config.DB

	param := c.Params("id")

	book := new(models.Book)

	if err := db.Find(&book, param); err != nil {
		c.Status(404).JSON(fiber.Map{"message": "Data is not found"})
	}

	db.Delete(&book, param)

	return c.Status(200).JSON(fiber.Map{"data": book, "message": "success delete the data"})

}
