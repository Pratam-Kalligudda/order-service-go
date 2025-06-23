package domain

import "time"

type Cart struct {
	ID        uint      `json:"id,omitempty"`
	UserID    uint      `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CartItem struct {
	ID          uint    `json:"id,omitempty"`
	CartID      uint    `json:"order_id,omitempty" gorm:""`
	ProductID   uint    `json:"product_id,omitempty"`
	ProductName string  `json:"product_name,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Cart        Cart    `gorm:"foreignKey:CartID; references:ID"`
}
