package helper

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Pratam-Kalligudda/order-service-go/internal/domain"
	"github.com/gofiber/fiber/v3"
)

type Auth struct {
	productServiceURL string
	secret            string
}

func NewAuth(sec, url string) Auth {
	return Auth{secret: sec}
}

func (a Auth) Authorize(c fiber.Ctx) error {
	log.Print("Authorized")
	return c.Next()
}

func (a Auth) GetProductDetails(productID uint) (domain.Product, error) {
	var product domain.Product
	url := a.productServiceURL + "/" + strconv.FormatUint(uint64(productID), 10)
	res, err := http.Get(url)
	if err != nil {
		return domain.Product{}, err
	}

	err = json.NewDecoder(res.Body).Decode(&product)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}
