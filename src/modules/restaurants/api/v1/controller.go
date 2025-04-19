package v1

import (
	"github.com/gofiber/fiber/v2"
	"goose/src/global"
	"goose/src/modules/restaurants/api/v1/dto"
)

func Create(c *fiber.Ctx) error {
	payload := new(dto.CreateRestaurantReq)
	err := c.BodyParser(payload)
	if err != nil {
		return err
	}
	res := createRestaurant(c, *payload)
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.CreateRestaurantRes]{
		Message: "Restaurant created successfully",
		Data:    res,
	})
}
