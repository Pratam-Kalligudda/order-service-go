package helper

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

type Auth struct {
	secret string
}

func NewAuth(sec string) Auth {
	return Auth{secret: sec}
}

func (a Auth) Authorize(c fiber.Ctx) error {
	log.Print("Authorized")
	return c.Next()
}
