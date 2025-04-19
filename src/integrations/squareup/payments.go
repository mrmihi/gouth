package squareup

import (
	"goose/src/modules/payments/api/v1/models"
)

func CreatePayment(payload models.Payment) {
	getClient().R().
		SetBody(payload).
		Post("/payments")

}
