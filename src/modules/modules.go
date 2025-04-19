package modules

import (
	"github.com/gofiber/fiber/v2"
	"goose/src/modules/orders"
	"goose/src/modules/payments"
	"goose/src/modules/restaurants"
	"goose/src/modules/system"
)

func New() *fiber.App {
	modules := fiber.New()

	modules.Mount("/system", system.New())

	modules.Mount("/restaurants", restaurants.New())

	modules.Mount("/orders", orders.New())

	modules.Mount("/payments", payments.New())

	return modules
}
