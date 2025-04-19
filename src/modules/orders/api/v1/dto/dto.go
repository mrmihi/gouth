package dto

import (
	"goose/src/modules/orders/api/v1/models"
)

type CreateOrderReq struct {
	OpenedAt string      `json:"opened_at"`
	IsClosed bool        `json:"is_closed"`
	Table    string      `json:"table"`
	Items    []OrderItem `json:"items"`
	Totals   OrderTotals `json:"totals"`
}

type OrderItem struct {
	Name      string     `json:"name"`
	Comment   string     `json:"comment"`
	UnitPrice int        `json:"unit_price"`
	Quantity  int        `json:"quantity"`
	Discounts []Discount `json:"discounts"`
	Modifiers []Modifier `json:"modifiers"`
	Amount    int        `json:"amount"`
}

type Discount struct {
	Name         string `json:"name"`
	IsPercentage bool   `json:"is_percentage"`
	Value        int    `json:"value"`
	Amount       int    `json:"amount"`
}

type Modifier struct {
	Name      string `json:"name"`
	UnitPrice int    `json:"unit_price"`
	Quantity  int    `json:"quantity"`
	Amount    int    `json:"amount"`
}

type OrderTotals struct {
	Discounts     int `json:"discounts"`
	Due           int `json:"due"`
	Tax           int `json:"tax"`
	ServiceCharge int `json:"service_charge"`
	Paid          int `json:"paid"`
	Tips          int `json:"tips"`
	Total         int `json:"total"`
}

type CreateOrderRes struct {
	Data models.Order `json:"data"`
}
