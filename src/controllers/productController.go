package controllers

import (
	"context"
	"encoding/json"
	"sort"
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

	go deleteCache("products_frontend")
	go deleteCache("products_backend")

	return c.JSON(fiber.Map{
		"data": product,
	})
}

func deleteCache(key string) {
	time.Sleep(5 * time.Second)
	database.Cache.Del(context.Background(), key)
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
		products = searchedProducts
	}

	if sortParams := c.Query("sort"); sortParams != "" {
		sortLower := strings.ToLower(sortParams)
		if sortLower == "asc" {
			sort.Slice(products, func(i, j int) bool {
				return products[i].Price < products[j].Price
			})
		} else if sortLower == "desc" {
			sort.Slice(products, func(i, j int) bool {
				return products[i].Price > products[j].Price
			})
		}

	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	total := len(products)
	perPage := 9
	var data []models.Product = products

	if total <= page*perPage && total >= (page-1)*perPage {
		data = products[(page-1)*perPage : total]
	} else if total >= page*perPage {
		data = products[(page-1)*perPage : perPage*page]
	} else {
		data = []models.Product{}
	}

	return c.JSON(fiber.Map{
		"data":      data,
		"total":     total,
		"page":      page,
		"last_page": total/perPage + 1,
	})
}
