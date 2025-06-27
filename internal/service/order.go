package service

import (
	"errors"

	"github.com/Pratam-Kalligudda/order-service-go/internal/domain"
	"github.com/Pratam-Kalligudda/order-service-go/internal/helper"
	"github.com/Pratam-Kalligudda/order-service-go/internal/repository"
)

type OrderService struct {
	repo        repository.OrderRepository
	cartService *CartService
	auth        helper.Auth
}

func NewOrderService(repo repository.OrderRepository, auth helper.Auth, cartService *CartService) *OrderService {
	return &OrderService{repo: repo, auth: auth, cartService: cartService}
}

func (svc *OrderService) CheckoutCartItems(userId uint) ([]domain.OrderItem, error) {
	cartItems, err := svc.cartService.GetCartItems(userId)
	if err != nil || len(cartItems) <= 0 {
		return nil, errors.New("cart items not found" + err.Error())
	}
	orderItems, err := svc.auth.MapCartItemsOrderItems(cartItems)
	if err != nil {
		return nil, err
	}
	totalAmount := 0.0
	for _, item := range orderItems {
		totalAmount += item.TotalAmount
	}
	order := domain.Order{
		UserID:      userId,
		TotalAmount: totalAmount,
		Status:      domain.PENDING,
	}
	order, err = svc.repo.CreateNewOrder(order)
	if err != nil {
		return nil, err
	}

	for i := range orderItems {
		orderItems[i].OrderID = order.ID
	}

	err = svc.repo.CheckoutCartItems(orderItems)
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}

func (svc *OrderService) GetOrders(userId uint) ([]domain.Order, error) {
	orders, err := svc.repo.GetOrders(userId)
	if err != nil {
		return nil, err
	}
	if len(orders) <= 0 {
		return nil, errors.New("no orders found")
	}
	return orders, nil
}

func (svc *OrderService) GetOrderByID(orderID uint) (domain.Order, error) {
	order, err := svc.repo.GetOrderByID(orderID)
	if err != nil {
		return domain.Order{}, err
	}
	if order == (domain.Order{}) {
		return domain.Order{}, errors.New("no order found")
	}
	return order, nil
}
