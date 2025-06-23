package domain

import "time"

const (
	PENDING   = "pending"
	CONFIRMED = "confirmed"
	CANCELED  = "canceled"
)

type Order struct {
	ID          uint      `json:"id,omitempty"`
	UserID      uint      `json:"user_id,omitempty"`
	TotalAmount float64   `json:"total_amount,omitempty"`
	Status      string    `json:"status,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type OrderItem struct {
	ID          uint    `json:"id,omitempty"`
	OrderID     uint    `json:"order_id,omitempty" gorm:""`
	ProductID   uint    `json:"product_id,omitempty"`
	ProductName string  `json:"product_name,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	Price       float64 `json:"price,omitempty"`
	TotalAmount float64 `json:"total_amount,omitempty"`
	Order       Order   `gorm:"foreignKey:OrderID; references:ID"`
}
