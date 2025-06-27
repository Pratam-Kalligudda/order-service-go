package repository

import (
	"gorm.io/gorm"

	"github.com/Pratam-Kalligudda/order-service-go/internal/domain"
)

type CartRepository interface {
	CreateCart(uint) (uint, error)
	GetCartItems(uint) ([]domain.CartItem, error)
	PostCartItem(domain.CartItem) (domain.CartItem, error)
	UpdateCartItem(uint, uint, int) (domain.CartItem, error)
	DeleteCartItem(uint, uint) error
	ClearCartItem(uint) error
	GetCartIDForUserID(uint) (uint, error)
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (ct *cartRepository) CreateCart(userId uint) (uint, error) {
	cart := domain.Cart{UserID: userId}
	err := ct.db.Create(&cart).Error
	return cart.ID, err
}

func (ct *cartRepository) GetCartItems(id uint) ([]domain.CartItem, error) {
	var cartItems []domain.CartItem
	err := ct.db.Model(&domain.CartItem{}).Find(&cartItems, "cart_id = ?", id).Error
	return cartItems, err
}

func (ct *cartRepository) PostCartItem(item domain.CartItem) (domain.CartItem, error) {
	err := ct.db.Create(&item).Error
	return item, err
}

func (ct *cartRepository) UpdateCartItem(cartId, productId uint, quantity int) (domain.CartItem, error) {
	var cartItem domain.CartItem
	err := ct.db.Model(&cartItem).Where("cart_id = ?", cartId).Where("product_id = ?", productId).Update("quantity", quantity).Error
	return cartItem, err
}

func (ct *cartRepository) DeleteCartItem(productId, cartId uint) error {
	err := ct.db.Where("cart_id = ?", cartId).Where("product_id = ?", productId).Delete(&domain.CartItem{}).Error
	return err
}

func (ct *cartRepository) ClearCartItem(cartId uint) error {
	err := ct.db.Where("cart_id = ?", cartId).Delete(&domain.CartItem{}).Error
	return err
}

func (ct *cartRepository) GetCartIDForUserID(id uint) (uint, error) {
	var cart domain.Cart
	err := ct.db.First(&cart, "id = ?", id).Error
	if err != nil {
		return 0, err
	}
	return cart.ID, nil
}
