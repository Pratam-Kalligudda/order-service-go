package handler

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"

	"github.com/Pratam-Kalligudda/order-service-go/internal/api/rest"
	"github.com/Pratam-Kalligudda/order-service-go/internal/repository"
	"github.com/Pratam-Kalligudda/order-service-go/internal/service"
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
	cartDB := repository.NewCartRepository(rh.DB)
	cartSvc := service.NewCartService(cartDB, rh.Auth)
	svc := service.NewOrderService(repo, rh.Auth, &cartSvc)
	handler := OrderHandler{svc: svc}

	orderRoutes := app.Group("/api/orders", rh.Auth.Authorize)
	orderRoutes.Get("/", handler.ListOrders)
	orderRoutes.Post("/checkout", handler.OrderCartItem)
	orderRoutes.Get("/:id", handler.GetOrderDetail)
}

func (h *OrderHandler) OrderCartItem(c fiber.Ctx) error {
	userId, ok := c.Locals("userId").(uint)
	if !ok || userId <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(Json{"error": "invalid userId"})
	}

	orderItems, err := h.svc.CheckoutCartItems(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Json{"error": "something went wrong : " + err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(Json{"message": "succesfully ordered cart item", "items": orderItems})
}

func (h *OrderHandler) ListOrders(c fiber.Ctx) error {
	userId, ok := c.Locals("userId").(uint)
	if !ok || userId <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(Json{"error": "invalid userId"})
	}

	orders, err := h.svc.GetOrders(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Json{"error": "something went wrong : " + err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(Json{"message": "succesfully listed orders", "orders": orders})
}

func (h *OrderHandler) GetOrderDetail(c fiber.Ctx) error {
	orderIdStr := c.Params("id")
	if orderIdStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(Json{"error": "invalid param"})
	}
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil || orderId <= 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(Json{"error": "invalid order id"})
	}

	order, err := h.svc.GetOrderByID(uint(orderId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Json{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(Json{"message": "succesfully got order details", "order": order})
}
