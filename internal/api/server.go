package api

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Pratam-Kalligudda/order-service-go/config"
	"github.com/Pratam-Kalligudda/order-service-go/internal/api/rest"
	"github.com/Pratam-Kalligudda/order-service-go/internal/api/rest/handler"
	"github.com/Pratam-Kalligudda/order-service-go/internal/domain"
	"github.com/Pratam-Kalligudda/order-service-go/internal/helper"
)

func SetupServer(config config.Config) {
	log.Println("reached 1")
	app := fiber.New()
	log.Println("reached 2")
	app.Use(logger.New())
	log.Println("reached 3")
	db, err := gorm.Open(postgres.Open(config.DNS), &gorm.Config{})
	if err != nil {
		log.Fatalf("error while connecting db : %v", err.Error())
	}

	log.Print("db connection succesfull")
	db.AutoMigrate(&domain.Cart{}, &domain.CartItem{})
	log.Println("url : " + config.PRODUCT_SERVICE_URL)
	httpHandler := rest.HTTPHandler{
		App:  app,
		DB:   db,
		Auth: helper.NewAuth(config.SECRET, config.PRODUCT_SERVICE_URL),
	}

	setupHandler(httpHandler)

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "order service is healthy"})
	})

	app.Listen(config.HOST)
}

func setupHandler(rh rest.HTTPHandler) {
	handler.SetupOrderHandler(rh)
	handler.SetupCartHandler(rh)
}
