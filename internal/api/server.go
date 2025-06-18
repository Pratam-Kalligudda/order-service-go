package api

import (
	"log"

	"github.com/Pratam-Kalligudda/order-service-go/config"
	"github.com/Pratam-Kalligudda/order-service-go/internal/api/rest"
	"github.com/Pratam-Kalligudda/order-service-go/internal/api/rest/handler"
	"github.com/Pratam-Kalligudda/order-service-go/internal/helper"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupServer(config config.Config) {
	app := fiber.New()
	app.Use(logger.New())

	db, err := gorm.Open(postgres.Open(config.DNS), &gorm.Config{})
	if err != nil {
		log.Fatalf("error while connecting db : %v", err.Error())
	}

	log.Print("db connection succesfull")
	db.AutoMigrate()

	httpHandler := rest.HTTPHandler{
		App:  app,
		DB:   db,
		Auth: helper.NewAuth(config.SECRET),
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
