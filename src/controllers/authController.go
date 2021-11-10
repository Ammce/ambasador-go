package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/Ammce/ambasador-go/src/database"
	"github.com/Ammce/ambasador-go/src/middlewares"
	"github.com/Ammce/ambasador-go/src/models"
	"github.com/gofiber/fiber/v2"
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

	user := models.User{
		FirstName:   data["first_name"],
		LastName:    data["last_name"],
		Email:       data["email"],
		IsAmbasador: strings.Contains(c.Path(), "/api/ambasador"),
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(&user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Bad credentials",
		})
	}

	if err := user.CompareHashAndPassword([]byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Bad credentials",
		})
	}

	scope := "admin"

	if strings.Contains(c.Path(), "/api/ambasador") {
		scope = "ambasador"
	}

	token, err := middlewares.GenerateToken(user.Id, scope)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Error generating tokenn",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func User(c *fiber.Ctx) error {
	var user models.User
	userId, _ := middlewares.GetUserId(c)
	dbErr := database.DB.Where("id = ?", userId).First(&user)

	if dbErr != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(fiber.Map{
			"message": "Error while accessing the user",
		})
	}

	if strings.Contains(c.Path(), "/api/ambasador") {
		ambasador := models.Ambasador(user)
		ambasador.CalculateRevenue(database.DB)
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"ambasador": user,
		})
	}

	admin := models.Admin(user)
	admin.CalculateRevenue(database.DB)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"admin": admin,
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	userId, _ := middlewares.GetUserId(c)
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		BaseModel: models.BaseModel{
			Id: userId,
		},
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	database.DB.Model(&user).Updates(&user)

	return c.JSON(fiber.Map{
		"user": user,
	})
}

func UpdatePassword(c *fiber.Ctx) error {
	userId, _ := middlewares.GetUserId(c)

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		BaseModel: models.BaseModel{
			Id: userId,
		},
	}

	user.SetPassword(data["password"])

	database.DB.Model(&user).Updates(&user)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
