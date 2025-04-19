package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goose/src/database"
	"goose/src/modules/payments/api/v1/dto"
	"time"
)

type Payment struct {
	IdempotencyKey string             `json:"idempotency_key" bson:"idempotency_key"`
	SourceId       string             `json:"source_id" bson:"source_id"`
	AmountMoney    Money              `json:"amount_money" bson:"amount_money"`
	CashDetails    CashPaymentDetails `json:"cash_details"`
	Autocomplete   bool               `json:"auto_complete" bson:"autocomplete"`
	LocationId     string             `json:"location_id" bson:"location_id"`
	OrderId        string             `json:"order_id" bson:"order_id"`
	TipMoney       Money              `json:"tip_money" bson:"tip_money"`
	UpdatedAt      string             `json:"updated_at" bson:"updated_at"`
	CreatedAt      string             `json:"created_at" bson:"created_at"`
}

type CashPaymentDetails struct {
	BuyerSuppliedMoney Money `json:"buyer_supplied_money"`
}

type Money struct {
	Amount   int    `json:"amount" bson:"amount"`
	Currency string `json:"currency" bson:"currency"`
}

func SyncIndexes() {
	database.UseDefault().Collection("payments").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "table", Value: -1}},
		Options: options.Index().SetUnique(true),
	})
}
func (payment Payment) WithDefaults() Payment {
	payment.CreatedAt = time.Now().Format(time.RFC3339)
	payment.UpdatedAt = time.Now().Format(time.RFC3339)
	return payment
}

func (payment *Payment) ToLocalPayment(p *dto.CreatePaymentReq) {
	payment.AmountMoney.Amount = int(p.BillAmount)
	payment.TipMoney.Amount = int(p.TipAmount)
	payment.IdempotencyKey = p.PaymentId
	payment.SourceId = "CASH"
	payment.OrderId = p.OrderId
	payment.CashDetails.BuyerSuppliedMoney.Amount = int((p.BillAmount + p.TipAmount))
	payment.AmountMoney.Currency = "USD"
	payment.TipMoney.Currency = "USD"
	payment.CashDetails.BuyerSuppliedMoney.Currency = "USD"
}
