package helper

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"

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
	log.Print("Authorized")
	tokenStr := c.GetReqHeaders()["Authorization"]
	if len(tokenStr) <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid token"})
	}

	tkn := strings.Split(tokenStr[0], " ")
	if len(tkn) < 2 || tkn[0] != "Bearer" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid token"})
	}

	usr, err := a.verifyToken(tkn[1])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid token : " + err.Error()})
	}

	c.Locals("userId", uint(usr.UserId))
	log.Printf("userId || Authorize: %v", usr.UserId)
	return c.Next()
}

func (a Auth) verifyToken(tokenStr string) (user, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(a.secret), nil
	})
	if err != nil || !token.Valid {
		return user{}, err
	}

	claims := token.Claims.(jwt.MapClaims)
	id, ok := claims["sub"].(float64)
	if !ok {
		return user{}, err
	}

	log.Printf("user_id || verifyToken: %v", id)
	user := user{
		UserId: id,
	}
	return user, nil
}

func (a Auth) GetProductDetails(productID uint) (domain.Product, error) {
	var response domain.ProductResponse
	url := a.productServiceURL + "/product/" + strconv.FormatUint(uint64(productID), 10)
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
