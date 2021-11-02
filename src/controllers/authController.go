package controllers

import (
	"net/http"

	"github.com/Ammce/ambasador-go/src/database"
	"github.com/Ammce/ambasador-go/src/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Password do not match",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

	user := models.User{
		FirstName:   data["first_name"],
		LastName:    data["last_name"],
		Email:       data["email"],
		IsAmbasador: false,
		Password:    password,
	}

	database.DB.Create(&user)

	return c.JSON(&user)
}
