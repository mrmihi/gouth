package v1

import (
	"goose/src/modules/payments/api/v1/models"
	"goose/src/utils"
)

var repository = utils.NewRepository[models.Payment]("payments")

func Repository() utils.Repository[models.Payment] {
	return repository
}
