package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goose/src/database"
	"time"
)

type Order struct {
	SquareUpOrderId string      `json:"squareup_order_id,omitempty" bson:"squareup_order_id,omitempty"`
	OpenedAt        string      `json:"opened_at,omitempty" bson:"opened_at,omitempty"`
	IsClosed        bool        `json:"is_closed,omitempty" bson:"is_closed,omitempty"`
	Table           string      `json:"table,omitempty" bson:"table,omitempty"`
	Items           []OrderItem `json:"items,omitempty" bson:"items,omitempty"`
	Totals          OrderTotals `json:"totals,omitempty" bson:"totals,omitempty"`
	CreatedAt       string      `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt       string      `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type OrderItem struct {
	Name      string     `json:"name,omitempty" bson:"name,omitempty"`
	Comment   string     `json:"comment,omitempty" bson:"comment,omitempty"`
	UnitPrice int        `json:"unit_price,omitempty" bson:"unit_price,omitempty"`
	Quantity  string     `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Discounts []Discount `json:"discounts,omitempty" bson:"discounts,omitempty"`
	Modifiers []Modifier `json:"modifiers,omitempty" bson:"modifiers,omitempty"`
	Amount    int        `json:"amount,omitempty" bson:"amount,omitempty"`
}

type Discount struct {
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	IsPercentage bool   `json:"is_percentage,omitempty" bson:"is_percentage,omitempty"`
	Value        int    `json:"value,omitempty" bson:"value,omitempty"`
	Amount       int    `json:"amount,omitempty" bson:"amount,omitempty"`
}

type Modifier struct {
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	UnitPrice int    `json:"unit_price,omitempty" bson:"unit_price,omitempty"`
	Quantity  string `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Amount    int    `json:"amount,omitempty" bson:"amount,omitempty"`
}

type OrderTotals struct {
	Discounts     int `json:"discounts,omitempty" bson:"discounts,omitempty"`
	Due           int `json:"due,omitempty" bson:"due,omitempty"`
	Tax           int `json:"tax,omitempty" bson:"tax,omitempty"`
	ServiceCharge int `json:"service_charge,omitempty" bson:"service_charge,omitempty"`
	Paid          int `json:"paid,omitempty" bson:"paid,omitempty"`
	Tips          int `json:"tips,omitempty" bson:"tips,omitempty"`
	Total         int `json:"total,omitempty" bson:"total,omitempty"`
}

func (order *Order) FromModel(o Order) {
	order.OpenedAt = o.OpenedAt
	order.IsClosed = o.IsClosed
	order.Table = o.Table

	order.Items = make([]OrderItem, len(o.Items))
	for i, item := range o.Items {
		order.Items[i] = OrderItem{
			Name:      item.Name,
			Comment:   item.Comment,
			UnitPrice: item.UnitPrice,
			Quantity:  item.Quantity,
			Amount:    item.Amount,
		}

		order.Items[i].Discounts = make([]Discount, len(item.Discounts))
		for j, discount := range item.Discounts {
			order.Items[i].Discounts[j] = Discount{
				Name:         discount.Name,
				IsPercentage: discount.IsPercentage,
				Value:        discount.Value,
				Amount:       discount.Amount,
			}
		}

		order.Items[i].Modifiers = make([]Modifier, len(item.Modifiers))
		for j, modifier := range item.Modifiers {
			order.Items[i].Modifiers[j] = Modifier{
				Name:      modifier.Name,
				UnitPrice: modifier.UnitPrice,
				Quantity:  modifier.Quantity,
				Amount:    modifier.Amount,
			}
		}
	}

	order.Totals = OrderTotals{
		Discounts:     o.Totals.Discounts,
		Due:           o.Totals.Due,
		Tax:           o.Totals.Tax,
		ServiceCharge: o.Totals.ServiceCharge,
		Paid:          o.Totals.Paid,
		Tips:          o.Totals.Tips,
		Total:         o.Totals.Total,
	}
}

func (order Order) WithDefaults() Order {
	order.CreatedAt = time.Now().Format(time.RFC3339)
	order.UpdatedAt = time.Now().Format(time.RFC3339)
	return order
}

func SyncIndexes() {
	database.UseDefault().Collection("orders").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "table", Value: -1}},
		Options: options.Index().SetUnique(true),
	})
}
