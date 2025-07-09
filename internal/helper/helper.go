package helper

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"

	"github.com/Pratam-Kalligudda/order-service-go/internal/domain"
)

type Auth struct {
	productServiceURL string
	secret            string
}

type user struct {
	UserId float64
}

func NewAuth(sec, url string) Auth {
	return Auth{secret: sec, productServiceURL: url}
}

func (a Auth) Authorize(c fiber.Ctx) error {
	userIdStr := c.GetReqHeaders()["X-User-Id"]
	if len(userIdStr) <= 0 || len(userIdStr[0]) <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid userId"})
	}
	userId, err := strconv.ParseUint(userIdStr[0], 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid userId : " + err.Error()})
	}
	c.Locals("userId", uint(userId))
	log.Printf("userId || Authorize: %v", userId)
	return c.Next()
}

func (a Auth) GetProductDetails(productID uint) (domain.Product, error) {
	var response domain.ProductResponse
	url := a.productServiceURL + "/api/products/" + strconv.FormatUint(uint64(productID), 10)
	res, err := http.Get(url)
	if err != nil {
		return domain.Product{}, err
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return domain.Product{}, err
	}
	return response.Product, nil
}

func (a Auth) MapCartItemsOrderItems(cartItems []domain.CartItem) ([]domain.OrderItem, error) {
	var orderItems []domain.OrderItem
	for _, cartItem := range cartItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductName: cartItem.ProductName,
			ProductID:   cartItem.ProductID,
			Quantity:    cartItem.Quantity,
			Price:       cartItem.Price,
			TotalAmount: (cartItem.Price * (float64(cartItem.Quantity))),
		})
	}
	if len(orderItems) <= 0 {
		return nil, errors.New("no cart item found")
	}
	return orderItems, nil
}
