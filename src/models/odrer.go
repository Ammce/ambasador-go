package models

type Order struct {
	BaseModel
	TransactionId  string      `json:"transaction_id" gorm:"null"`
	UserId         uint        `json:"user_id"`
	Code           string      `json:"code"`
	AmbasadorEmail string      `json:"ambasador_email"`
	FirstName      string      `json:"-"`
	LastName       string      `json:"-"`
	Name           string      `json:"name" gorm:"-"`
	Email          string      `json:"email"`
	Address        string      `json:"address" gorm:"null"`
	City           string      `json:"city" gorm:"null"`
	Country        string      `json:"country" gorm:"null"`
	Zip            string      `json:"zip" gorm:"null"`
	Complete       bool        `json:"-" gorm:"default:false"`
	Total          float64     `json:"total" gorm:"-"`
	OrderItem      []OrderItem `json:"order_item" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	BaseModel
	OrderId          uint    `json:"order_id"`
	ProductTitle     string  `json:"product_title"`
	Price            float64 `json:"price"`
	Quantity         uint    `json:"quantity"`
	AdminRevenue     float64 `json:"admin_revenue"`
	AmbasadorRevenue float64 `json:"ambasador_revenue"`
}

func (o *Order) GetFullName() string {
	return o.FirstName + " " + o.LastName
}

func (o *Order) GetTotal() float64 {
	var total float64 = 0
	for _, orderItem := range o.OrderItem {
		total += orderItem.Price * float64(orderItem.Quantity)
	}
	return total
}
