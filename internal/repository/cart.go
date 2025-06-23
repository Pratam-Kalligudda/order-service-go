package repository

import (
	"github.com/Pratam-Kalligudda/order-service-go/internal/domain"
	"gorm.io/gorm"
)

type CartRepository interface {
	GetCartItems(uint) ([]domain.CartItem, error)
	PostCartItem(uint) (domain.CartItem, error)
	UpdateCartItem(uint) (domain.CartItem, error)
	DeleteCartItem(uint) (domain.CartItem, error)
	ClearCartItem(uint) error
	GetCartIDForUserID(uint) (uint, error)
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}
func (ct *cartRepository) GetCartItems(uint) ([]domain.CartItem, error) {
	return nil, nil
}
func (ct *cartRepository) PostCartItem(uint) (domain.CartItem, error) {
	return domain.CartItem{}, nil
}
func (ct *cartRepository) UpdateCartItem(uint) (domain.CartItem, error) {
	return domain.CartItem{}, nil
}
func (ct *cartRepository) DeleteCartItem(uint) (domain.CartItem, error) {
	return domain.CartItem{}, nil
}
func (ct *cartRepository) ClearCartItem(uint) error {
	return nil
}
func (ct *cartRepository) GetCartIDForUserID(uint) (uint, error) {
	return 0, nil
}
