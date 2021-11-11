package controllers

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

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

	product := models.Product{
		BaseModel: models.BaseModel{
			Id: uint(id),
		},
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(&product)

	return c.JSON(fiber.Map{
		"data": product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		BaseModel: models.BaseModel{
			Id: uint(id),
		},
	}

	database.DB.Delete(&product)
	return nil
}

func ProductsFrontend(c *fiber.Ctx) error {
	var products []models.Product

	var ctx = context.Background()

	result, err := database.Cache.Get(ctx, "products_frontend").Result()

	if err != nil {
		database.DB.Find(&products)
		bytes, errMarsh := json.Marshal(products)

		if errMarsh != nil {
			panic(errMarsh)
		}

		if errKey := database.Cache.Set(ctx, "products_frontend", bytes, 30*time.Minute).Err(); errKey != nil {
			panic(errKey)
		}

	} else {
		json.Unmarshal([]byte(result), &products)
	}

	return c.JSON(products)
}

func ProductsBackend(c *fiber.Ctx) error {
	var products []models.Product

	var ctx = context.Background()

	result, err := database.Cache.Get(ctx, "products_backend").Result()

	if err != nil {
		database.DB.Find(&products)
		bytes, errMarsh := json.Marshal(products)

		if errMarsh != nil {
			panic(errMarsh)
		}

		database.Cache.Set(ctx, "products_backend", bytes, 30*time.Minute)

	} else {
		json.Unmarshal([]byte(result), &products)
	}

	var searchedProducts []models.Product

	if s := c.Query("s"); s != "" {
		for _, product := range products {
			if strings.Contains(product.Title, s) || strings.Contains(product.Description, s) {
				searchedProducts = append(searchedProducts, product)
			}
		}
		return c.JSON(searchedProducts)
	}

	return c.JSON(products)
}
