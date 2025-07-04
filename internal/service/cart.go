package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Pratam-Kalligudda/order-service-go/internal/domain"
	"github.com/Pratam-Kalligudda/order-service-go/internal/dto"
	"github.com/Pratam-Kalligudda/order-service-go/internal/helper"
	"github.com/Pratam-Kalligudda/order-service-go/internal/repository"
)

type CartService struct {
	repo repository.CartRepository
	auth helper.Auth
}

func NewCartService(repo repository.CartRepository, auth helper.Auth) CartService {
	return CartService{repo: repo, auth: auth}
}

func (s *CartService) GetCartItems(userId uint) ([]domain.CartItem, error) {
	cartId, err := s.repo.GetCartIDForUserID(userId)
	if err != nil {
		return nil, err
	}
	cartItems, err := s.repo.GetCartItems(cartId)
	if err != nil {
		return nil, err
	}

	if len(cartItems) <= 0 {
		return nil, errors.New("no cart items found")
	}

	return cartItems, nil
}

func (s *CartService) AddItemToCart(userId uint, ctItem dto.AddUpdateProduct) (domain.CartItem, error) {
	product, err := s.auth.GetProductDetails(ctItem.ProductID)
	if err != nil {
		return domain.CartItem{}, err
	}

	cartId, err := s.repo.GetCartIDForUserID(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cartId, err = s.repo.CreateCart(userId)
			if err != nil {
				return domain.CartItem{}, err
			}
		} else {
			return domain.CartItem{}, err
		}
	}

	item := domain.CartItem{
		CartID:      cartId,
		ProductID:   product.ID,
		ProductName: product.Name,
		Price:       product.Price,
		Quantity:    ctItem.Quantity,
	}

	item, err = s.repo.PostCartItem(item)
	if err != nil {
		return domain.CartItem{}, err
	}

	return item, nil
}

func (s *CartService) UpdateCartItem(userId uint, ctItem dto.AddUpdateProduct) error {
	cartId, err := s.repo.GetCartIDForUserID(userId)
	if err != nil {
		return err
	}
	_, err = s.repo.UpdateCartItem(cartId, ctItem.ProductID, ctItem.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func (s *CartService) RemoveCartItem(userId, productId uint) error {
	cartId, err := s.repo.GetCartIDForUserID(userId)
	if err != nil {
		return err
	}

	err = s.repo.DeleteCartItem(productId, cartId)
	if err != nil {
		return err
	}

	return nil
}

func (s *CartService) ClearCart(userId uint) error {
	cartId, err := s.repo.GetCartIDForUserID(userId)
	if err != nil {
		return err
	}

	err = s.repo.ClearCartItem(cartId)
	if err != nil {
		return err
	}

	return nil
}
