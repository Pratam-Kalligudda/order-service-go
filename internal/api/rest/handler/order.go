package handler

import (
	"log"

	"github.com/Pratam-Kalligudda/order-service-go/internal/api/rest"
	"github.com/Pratam-Kalligudda/order-service-go/internal/repository"
	"github.com/Pratam-Kalligudda/order-service-go/internal/service"
	"github.com/gofiber/fiber/v3"
)

type Json fiber.Map

type OrderHandler struct {
	svc *service.OrderService
}

// func NewOrderHandler(service.OrderService) OrderHandler {
// 	return
// }

func SetupOrderHandler(rh rest.HTTPHandler) {
	app := rh.App
	log.Print("order handler intialized")
	repo := repository.NewOrderRepository(rh.DB)
	svc := service.NewOrderService(repo, rh.Auth)
	handler := OrderHandler{svc: svc}

	orderRoutes := app.Group("/order", rh.Auth.Authorize)
	orderRoutes.Get("/", handler.ListOrders)
	orderRoutes.Post("/checkout", handler.OrderCartItem)
	orderRoutes.Get("/:id", handler.GetOrderDetail)

}

func (h *OrderHandler) OrderCartItem(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(Json{"message": "succesfully ordered cart item"})
}

func (h *OrderHandler) ListOrders(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(Json{"message": "succesfully listed orders"})
}

func (h *OrderHandler) GetOrderDetail(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(Json{"message": "succesfully got order details"})
}
