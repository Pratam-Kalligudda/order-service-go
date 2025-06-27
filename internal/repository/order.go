package repository

import (
	"gorm.io/gorm"

	"github.com/Pratam-Kalligudda/order-service-go/internal/domain"
)

type OrderRepository interface {
	CheckoutCartItems([]domain.OrderItem) error
	GetOrders(uint) ([]domain.Order, error)
	GetOrderByID(uint) (domain.Order, error)
	CreateNewOrder(domain.Order) (domain.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (oR *orderRepository) CheckoutCartItems(items []domain.OrderItem) error {
	err := oR.db.Create(&items).Error
	return err
}

func (oR *orderRepository) GetOrders(userId uint) ([]domain.Order, error) {
	var items []domain.Order
	err := oR.db.Find(&items, "user_id = ?", userId).Error
	return items, err
}

func (oR *orderRepository) GetOrderByID(orderID uint) (domain.Order, error) {
	var item domain.Order
	err := oR.db.Find(&item, orderID).Error
	return item, err
}

func (oR *orderRepository) CreateNewOrder(order domain.Order) (domain.Order, error) {
	err := oR.db.Create(&order).Error
	return order, err
}
