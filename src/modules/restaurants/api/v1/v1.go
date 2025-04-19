package v1

import (
	m "goose/src/middleware"
	"goose/src/modules/restaurants/api/v1/dto"
	v1 "goose/src/modules/restaurants/api/v1/models"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	v1.SyncIndexes()
	v1 := fiber.New()
	v1.Post("/", m.Validate[dto.CreateRestaurantReq](m.Body), Create)
	return v1
}
