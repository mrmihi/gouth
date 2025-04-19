package dto

import (
	"goose/src/modules/orders/api/v1/models"
	"strconv"
	"time"
)

type Error struct {
	Code     string `json:"code"`
	Detail   string `json:"detail"`
	Category string `json:"category"`
}

type Errors struct {
	Errors []Error `json:"errors"`
}
type GetSquareOrderByTableRes struct {
	Orders []SquareUpOrder `json:"orders"`
}
type CreateSquareOrderRes struct {
	Errors []Error       `json:"errors"`
	Order  SquareUpOrder `json:"order"`
}

type SquareUpOrder struct {
	ID                      string                 `json:"id"`
	LocationID              string                 `json:"location_id"`
	LineItems               []SquareUpLineItem     `json:"line_items"`
	CreatedAt               time.Time              `json:"created_at"`
	UpdatedAt               time.Time              `json:"updated_at"`
	State                   string                 `json:"state"`
	Version                 int                    `json:"version"`
	ReferenceID             string                 `json:"reference_id"`
	TotalTaxMoney           SquareUpMoney          `json:"total_tax_money"`
	TotalDiscountMoney      SquareUpMoney          `json:"total_discount_money"`
	TotalTipMoney           SquareUpMoney          `json:"total_tip_money"`
	TotalMoney              SquareUpMoney          `json:"total_money"`
	TotalServiceChargeMoney SquareUpMoney          `json:"total_service_charge_money"`
	NetAmounts              SquareUpNetAmounts     `json:"net_amounts"`
	Source                  SquareUpSource         `json:"source"`
	PricingOptions          SquareUpPricingOptions `json:"pricing_options"`
	TicketName              string                 `json:"ticket_name"`
	NetAmountDueMoney       SquareUpMoney          `json:"net_amount_due_money"`
}

type SquareUpLineItem struct {
	UID                      string             `json:"uid"`
	Quantity                 int                `json:"quantity"`
	Name                     string             `json:"name"`
	BasePriceMoney           SquareUpMoney      `json:"base_price_money"`
	Modifiers                []SquareUpModifier `json:"modifiers"`
	GrossSalesMoney          SquareUpMoney      `json:"gross_sales_money"`
	TotalTaxMoney            SquareUpMoney      `json:"total_tax_money"`
	TotalDiscountMoney       SquareUpMoney      `json:"total_discount_money"`
	TotalMoney               SquareUpMoney      `json:"total_money"`
	VariationTotalPriceMoney SquareUpMoney      `json:"variation_total_price_money"`
	ItemType                 string             `json:"item_type"`
	TotalServiceChargeMoney  SquareUpMoney      `json:"total_service_charge_money"`
}

type SquareUpModifier struct {
	UID             string        `json:"uid"`
	BasePriceMoney  SquareUpMoney `json:"base_price_money"`
	TotalPriceMoney SquareUpMoney `json:"total_price_money"`
	Name            string        `json:"name"`
	Quantity        int           `json:"quantity"`
}

type SquareUpMoney struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type SquareUpNetAmounts struct {
	TotalMoney         SquareUpMoney `json:"total_money"`
	TaxMoney           SquareUpMoney `json:"tax_money"`
	DiscountMoney      SquareUpMoney `json:"discount_money"`
	TipMoney           SquareUpMoney `json:"tip_money"`
	ServiceChargeMoney SquareUpMoney `json:"service_charge_money"`
}

type SquareUpSource struct {
	Name string `json:"name"`
}

type SquareUpPricingOptions struct {
	AutoApplyDiscounts bool `json:"auto_apply_discounts"`
	AutoApplyTaxes     bool `json:"auto_apply_taxes"`
}

func (res SquareUpOrder) ToOrder() models.Order {
	order := models.Order{
		SquareUpOrderId: res.ID,
		OpenedAt:        res.CreatedAt.Format(time.RFC3339),
		IsClosed:        res.State != "OPEN", // Assuming Opened if state is OPEN
		Table:           res.Source.Name,     // Assuming Source name is table
		CreatedAt:       res.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       res.UpdatedAt.Format(time.RFC3339),
	}

	// Map order items
	var items []models.OrderItem
	for _, item := range res.LineItems {
		orderItem := models.OrderItem{
			Name:      item.Name,
			UnitPrice: item.BasePriceMoney.Amount,
			Quantity:  strconv.Itoa(item.Quantity),
			Amount:    item.TotalMoney.Amount,
		}

		// Map modifiers
		var modifiers []models.Modifier
		for _, modifier := range item.Modifiers {
			modifiers = append(modifiers, models.Modifier{
				Name:      modifier.Name,
				UnitPrice: modifier.BasePriceMoney.Amount,
				Quantity:  strconv.Itoa(modifier.Quantity),
				Amount:    modifier.TotalPriceMoney.Amount,
			})
		}
		orderItem.Modifiers = modifiers

		items = append(items, orderItem)
	}
	order.Items = items

	// Map order totals
	orderTotals := models.OrderTotals{
		Discounts:     res.TotalDiscountMoney.Amount,
		Due:           res.NetAmountDueMoney.Amount,
		Tax:           res.TotalTaxMoney.Amount,
		ServiceCharge: res.TotalServiceChargeMoney.Amount,
		Paid:          res.TotalMoney.Amount,
		Tips:          res.TotalTipMoney.Amount,
		Total:         res.TotalMoney.Amount, // Assuming this is the total amount paid
	}
	order.Totals = orderTotals

	return order
}
