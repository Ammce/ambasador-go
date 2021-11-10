package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Email       string   `json:"email" gorm:"unique"`
	Password    []byte   `json:"-"`
	IsAmbasador bool     `json:"-"`
	Revenue     *float64 `json:"revenue, omitempty" gorm:"-"`
}

func (u *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	u.Password = hashedPassword
}

func (u *User) CompareHashAndPassword(password []byte) error {
	if err := bcrypt.CompareHashAndPassword(u.Password, password); err != nil {
		return err
	}
	return nil
}

type Admin User

type Ambasador User

func (a Admin) CalculateRevenue(db *gorm.DB) {
	var orders []Order
	db.Preload("OrderItems").Find(&orders, &Order{
		UserId:   a.Id,
		Complete: true,
	})

	var revenue float64 = 0

	for _, order := range orders {
		for _, orderItem := range order.OrderItem {
			revenue += orderItem.AdminRevenue
		}
	}

	a.Revenue = &revenue
}

func (a Ambasador) CalculateRevenue(db *gorm.DB) {
	var orders []Order
	db.Preload("OrderItems").Find(&orders, &Order{
		UserId:   a.Id,
		Complete: true,
	})

	var revenue float64 = 0

	for _, order := range orders {
		for _, orderItem := range order.OrderItem {
			revenue += orderItem.AmbasadorRevenue
		}
	}

	a.Revenue = &revenue
}
