package service

import (
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
