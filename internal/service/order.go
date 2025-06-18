package service

import (
	"github.com/Pratam-Kalligudda/order-service-go/internal/helper"
	"github.com/Pratam-Kalligudda/order-service-go/internal/repository"
)

type OrderService struct {
	repo repository.OrderRepository
	auth helper.Auth
}

func NewOrderService(repo repository.OrderRepository, auth helper.Auth) *OrderService {
	return &OrderService{repo: repo, auth: auth}
}
