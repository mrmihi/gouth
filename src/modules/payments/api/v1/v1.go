package v1

import (
	"github.com/gofiber/fiber/v2"
	m "goose/src/middleware"
	"goose/src/modules/payments/api/v1/dto"
)

func New() *fiber.App {

	v1 := fiber.New()
	v1.Post("/", m.Validate[dto.CreatePaymentReq](m.Body), Create)
	return v1
}
