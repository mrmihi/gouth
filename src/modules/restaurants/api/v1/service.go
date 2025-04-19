package v1

import (
	"github.com/google/uuid"
	"goose/src/modules/restaurants/api/v1/dto"
	"goose/src/modules/restaurants/api/v1/models"
	"goose/src/modules/restaurants/api/v1/repository"
	"goose/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sethvargo/go-password/password"
)

func createRestaurant(c *fiber.Ctx, payload dto.CreateRestaurantReq) *dto.CreateRestaurantRes {
	log.Info("Creating restaurant within system - ", payload.Email)
	generatedPassword, _ := password.Generate(8, 2, 1, false, false)
	generatedAPIKey := uuid.New().String()
	insertedID := repository.GetInstance().Create(models.Restaurant{
		Email:    payload.Email,
		Name:     payload.Name,
		Password: utils.HashStr(generatedPassword),
		APIKey:   utils.HashStr(generatedAPIKey),
	}.WithDefaults())
	return &dto.CreateRestaurantRes{
		ID:       insertedID,
		Name:     payload.Name,
		Email:    payload.Email,
		Password: generatedPassword,
		APIKey:   generatedAPIKey,
	}
}
