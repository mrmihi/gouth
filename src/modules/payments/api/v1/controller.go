package v1

import (
	"github.com/gofiber/fiber/v2"
	"goose/src/global"
	"goose/src/modules/payments/api/v1/dto"
)

func Create(c *fiber.Ctx) error {
	payload := new(dto.CreatePaymentReq)
	err := c.BodyParser(payload)
	if err != nil {
		return err
	}
	res := createPaymentService(c, *payload)
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.CreatePaymentRes]{
		Message: "Payment created successfully",
		Data:    res,
	})
}
