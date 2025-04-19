package restaurants

import (
	"github.com/gofiber/fiber/v2"
	"goose/src/modules/restaurants/api/v1"
)

func New() *fiber.App {
	restaurants := fiber.New()
	//restaurants.All("/*", middleware.AdminProtect)
	restaurants.Mount("/v1", v1.New())
	return restaurants
}
