package controllers

import (
	"strconv"

	"github.com/Ammce/ambasador-go/src/database"
	"github.com/Ammce/ambasador-go/src/models"
	"github.com/gofiber/fiber/v2"
)

func Products(c *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"data": products,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)

	return c.JSON(fiber.Map{
		"data": product,
	})
}

func Product(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var product models.Product

	product.Id = uint(id)

	database.DB.Find(&product)

	return c.JSON(fiber.Map{
		"data": product,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	priceFloat, _ := strconv.ParseFloat(data["price"], 64)

	product := models.Product{
		Id:          uint(id),
		Title:       data["title"],
		Description: data["description"],
		Price:       priceFloat,
		Image:       data["image"],
	}

	database.DB.Find(&product)

	return c.JSON(fiber.Map{
		"data": product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Delete(&product)
	return nil
}
