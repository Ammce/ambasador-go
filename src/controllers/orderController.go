package controllers

import (
	"github.com/Ammce/ambasador-go/src/database"
	"github.com/Ammce/ambasador-go/src/models"
	"github.com/gofiber/fiber/v2"
)

func Orders(c *fiber.Ctx) error {
	var orders []models.Order

	database.DB.Preload("OrderItem").Find(&orders)

	for index, order := range orders {
		orders[index].Name = order.GetFullName()
		orders[index].Total = order.GetTotal()
	}

	return c.JSON(orders)
}
