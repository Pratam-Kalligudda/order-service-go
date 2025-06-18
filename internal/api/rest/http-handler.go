package rest

import (
	"github.com/Pratam-Kalligudda/order-service-go/internal/helper"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type HTTPHandler struct {
	App  *fiber.App
	DB   *gorm.DB
	Auth helper.Auth
}
