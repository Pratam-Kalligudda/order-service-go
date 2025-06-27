package handler

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"

	"github.com/Pratam-Kalligudda/order-service-go/internal/api/rest"
	"github.com/Pratam-Kalligudda/order-service-go/internal/dto"
	"github.com/Pratam-Kalligudda/order-service-go/internal/repository"
	"github.com/Pratam-Kalligudda/order-service-go/internal/service"
)

type CartHandler struct {
	svc service.CartService
}

func SetupCartHandler(rh rest.HTTPHandler) {
	log.Print("cart handler intiazlied")
	app := rh.App
	repo := repository.NewCartRepository(rh.DB)
	svc := service.NewCartService(repo, rh.Auth)
	handler := CartHandler{svc}

	pvtRoutes := app.Group("/cart", rh.Auth.Authorize)
	pvtRoutes.Get("/", handler.GetCartItems)
	pvtRoutes.Post("/item/add", handler.AddItemToCart)
	pvtRoutes.Put("/item/update", handler.UpdateCartItem)
	pvtRoutes.Delete("/item/:id", handler.RemoveCartItem)
	pvtRoutes.Delete("/clear", handler.ClearCart)
}

func (h *CartHandler) GetCartItems(c fiber.Ctx) error {
	userId, ok := c.Locals("userId").(uint)
	if !ok || userId <= 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(Json{"error": "invalid userId : " + fmt.Sprint(userId)})
	}

	cartItems, err := h.svc.GetCartItems(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Json{"error": err})
	}

	return c.Status(fiber.StatusOK).JSON(Json{"message": "succesfully got cart items", "items": cartItems})
}

func (h *CartHandler) AddItemToCart(c fiber.Ctx) error {
	userId, ok := c.Locals("userId").(uint)
	if !ok || userId <= 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(Json{"error": "invalid userId : " + fmt.Sprint(userId)})
	}

	var dto dto.AddUpdateProduct
	err := c.Bind().Body(&dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Json{"error": err})
	}

	item, err := h.svc.AddItemToCart(userId, dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Json{"error": err})
	}
	return c.Status(fiber.StatusOK).JSON(Json{"message": "succesfully added item to cart", "item": item})
}

func (h *CartHandler) UpdateCartItem(c fiber.Ctx) error {
	userId, ok := c.Locals("userId").(uint)
	if !ok || userId <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(Json{"error": "invalid userId : " + fmt.Sprintf("type : %T value : %v", userId, userId)})
	}

	var dto dto.AddUpdateProduct
	err := c.Bind().Body(&dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Json{"error": "invalid body :" + err.Error()})
	}

	err = h.svc.UpdateCartItem(userId, dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Json{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(Json{"message": "succesfully updated cart item"})
}

func (h *CartHandler) RemoveCartItem(c fiber.Ctx) error {
	userId, ok := c.Locals("userId").(uint)
	if !ok || userId <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(Json{"error": "invalid userId : " + fmt.Sprintf("type : %T value : %v", userId, userId)})
	}

	productIdStr := c.Params("id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Json{"error": "invalid param : " + err.Error()})
	}

	err = h.svc.RemoveCartItem(userId, uint(productId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Json{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(Json{"message": "succesfully removed cart item"})
}

func (h *CartHandler) ClearCart(c fiber.Ctx) error {
	userId, ok := c.Locals("userId").(uint)
	if !ok || userId <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(Json{"error": "invalid userId : " + fmt.Sprintf("type : %T value : %v", userId, userId)})
	}

	err := h.svc.ClearCart(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Json{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(Json{"message": "succesfully cleared cart"})
}
