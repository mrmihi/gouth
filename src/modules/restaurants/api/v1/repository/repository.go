package repository

import (
	"goose/src/modules/restaurants/api/v1/models"
	"goose/src/utils"
)

var instance = utils.NewRepository[models.Restaurant]("restaurants")

func GetInstance() utils.Repository[models.Restaurant] {
	return instance
}
