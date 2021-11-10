package main

import (
	"math/rand"

	"github.com/Ammce/ambasador-go/src/database"
	"github.com/Ammce/ambasador-go/src/models"
	"github.com/bxcodec/faker/v3"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		var orderItems []models.OrderItem

		for j := 0; j < rand.Intn(5); j++ {
			price := float64(rand.Intn(90) + 10)
			quantity := uint(rand.Intn(5))
			orderItems = append(orderItems, models.OrderItem{
				ProductTitle:     faker.Username(),
				Price:            price,
				Quantity:         quantity,
				AdminRevenue:     0.9 * price * float64(quantity),
				AmbasadorRevenue: 0.1 * price * float64(quantity),
			})
		}
		database.DB.Create(&models.Order{
			UserId:         uint(rand.Intn(30) + 1),
			Code:           faker.Username(),
			AmbasadorEmail: faker.Email(),
			FirstName:      faker.FirstName(),
			LastName:       faker.LastName(),
			Email:          faker.Email(),
			Complete:       true,
			OrderItem:      orderItems,
		})
	}
}
