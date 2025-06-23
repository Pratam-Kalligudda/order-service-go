package repository

import (
	"github.com/Pratam-Kalligudda/order-service-go/internal/domain"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CheckoutCartItems(any) error
	GetOrders(uint) ([]domain.OrderItem, error)
	GetOrderByID(uint) (domain.OrderItem, error)
	GetOrderIDForUserID(uint) ([]uint, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}
func (od *orderRepository) CheckoutCartItems(any) error {
	return nil
}
func (od *orderRepository) GetOrders(uint) ([]domain.OrderItem, error) {
	return nil, nil
}
func (od *orderRepository) GetOrderByID(uint) (domain.OrderItem, error) {
	return domain.OrderItem{}, nil
}
func (od *orderRepository) GetOrderIDForUserID(uint) ([]uint, error) {
	return nil, nil
}
