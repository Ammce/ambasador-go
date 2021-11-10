package controllers

import (
	"github.com/Ammce/ambasador-go/src/database"
	"github.com/Ammce/ambasador-go/src/models"
	"github.com/gofiber/fiber/v2"
)

func GetAmbasadors(c *fiber.Ctx) error {
	var ambasadors []models.User

	database.DB.Where("is_ambasador = true").Find(&ambasadors)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"data": ambasadors,
	})

}
