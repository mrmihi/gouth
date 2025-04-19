package system

import (
	"github.com/gofiber/fiber/v2"
	"goose/src/global"
	"goose/src/modules/system/dto"
)

func Health(c *fiber.Ctx) error {
	return c.JSON(global.Response[*interface{}]{
		Message: "Goose up and running",
	})
}

func Memory(c *fiber.Ctx) error {
	return c.JSON(global.Response[dto.MemStats]{
		Message: "Memory usage retrieved",
		Data:    GetMemoryUsage(),
	})
}
