package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"goose/src/modules/restaurants/api/v1/repository"
	"goose/src/utils"
)

func RestaurantProtect(ctx *fiber.Ctx) error {
	token := ctx.Get("X-API-Key")
	email := ctx.Get("X-User-Email")
	if token == "" {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Missing API Key"))
	}
	restaurant := repository.GetInstance().FindOne(primitive.M{"email": email})
	if restaurant == nil {
		panic(fiber.NewError(fiber.StatusUnauthorized, "No restaurant with the given email"))
	}
	if !utils.CompareStrHash(token, restaurant.APIKey) {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Invalid API Key"))
	}

	ctx.Locals("restaurant", restaurant)
	return ctx.Next()
}
