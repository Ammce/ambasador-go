package controllers

import (
	"strconv"

	"github.com/Ammce/ambasador-go/src/database"
	"github.com/Ammce/ambasador-go/src/models"
	"github.com/gofiber/fiber/v2"
)

func GetUserLinks(c *fiber.Ctx) error {
	userId, _ := strconv.Atoi(c.Params("id"))
	var links []models.Link

	database.DB.Where("user_id = ?", userId).Find(&links)

	for i, link := range links {
		var orders []models.Order
		database.DB.Where("code = ? and complete = true", link.Code).Find(&orders)

		links[i].Orders = orders
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"data": links,
	})
}
