package v1

import (
	"goose/src/modules/orders/api/v1/models"
	"goose/src/utils"
)

var repository = utils.NewRepository[models.Order]("orders")

func Repository() utils.Repository[models.Order] {
	return repository
}
