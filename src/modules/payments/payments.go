package payments

import (
	"github.com/gofiber/fiber/v2"
	"goose/src/modules/payments/api/v1"
)

func New() *fiber.App {
	payments := fiber.New()
	payments.Mount("/v1", v1.New())
	return payments
}
