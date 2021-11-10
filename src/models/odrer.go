package models

type Order struct {
	BaseModel
	TransactionId  string      `json:"transaction_id" gorm:"null"`
	UserId         uint        `json:"user_id"`
	Code           string      `json:"code"`
	AmbasadorEmail string      `json:"ambasador_email"`
	FirstName      string      `json:"first_name"`
	LastName       string      `json:"last_name"`
	Email          string      `json:"email"`
	Address        string      `json:"address" gorm:"null"`
	City           string      `json:"city" gorm:"null"`
	Country        string      `json:"country" gorm:"null"`
	Zip            string      `json:"zip" gorm:"null"`
	Complete       bool        `json:"complete" gorm:"default:false"`
	OrderItem      []OrderItem `json:"order_item" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	BaseModel
	OrderId          uint    `json:"order_id"`
	ProductTitle     uint    `json:"product_title"`
	Price            float64 `json:"price"`
	Quantity         uint    `json:"quantity"`
	AdminRevenue     float64 `json:"admin_revenue"`
	AmbasadorRevenue float64 `json:"ambasador_revenue"`
}