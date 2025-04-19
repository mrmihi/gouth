package v1

import (
	"github.com/gofiber/fiber/v2"
	m "goose/src/middleware"
	"goose/src/modules/orders/api/v1/dto"
)

func New() *fiber.App {
	//v1.SyncIndexes()
	v1 := fiber.New()
	v1.Post("/", m.Validate[dto.CreateOrderReq](m.Body), Create)
	v1.Get("/:orderID", GetById)
	v1.Get("/table/:tableID", GetByTable)
	return v1
}
