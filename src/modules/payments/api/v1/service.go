package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"goose/src/integrations/squareup"
	"goose/src/modules/payments/api/v1/dto"
	"goose/src/modules/payments/api/v1/models"
)

func createPaymentService(c *fiber.Ctx, payload dto.CreatePaymentReq) *dto.CreatePaymentRes {
	log.Info("Creating payment within system")
	var payment models.Payment
	payment.ToLocalPayment(&payload)
	squareup.CreatePayment(payment)
	repository.Create(payment.WithDefaults())
	return &dto.CreatePaymentRes{}
}
