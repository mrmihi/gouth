package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"goose/src/integrations/squareup"
	"goose/src/modules/orders/api/v1/dto"
	"goose/src/modules/orders/api/v1/models"
	rest "goose/src/modules/restaurants/api/v1/models"
	"strconv"
)

// creates an order in SquareUp and the Database
func createOrderService(c *fiber.Ctx, payload dto.CreateOrderReq) *dto.CreateOrderRes {
	log.Info("Creating order within system")
	var order models.Order
	order = mapDTOToModel(payload)
	squareupOrderId := squareup.CreateOrder(order)
	order.SquareUpOrderId = squareupOrderId
	repository.Create(order.WithDefaults())
	return &dto.CreateOrderRes{
		Data: order,
	}
}

// retrieves an order by its ID from SquareUp and updates the Database
func getOrderServiceByID(c *fiber.Ctx, orderID string) *dto.CreateOrderRes {
	log.Info("Getting the order by ID within system")
	var order models.Order
	order = squareup.GetOrderById(orderID)
	query := primitive.M{"square_up_order_id": order.SquareUpOrderId}
	repository.UpdateBy(query, order)
	return &dto.CreateOrderRes{
		Data: order,
	}
}

// retrieves an order by its table from SquareUp and updates the Database
func getOrderServiceByTable(c *fiber.Ctx, tableID string) *dto.CreateOrderRes {
	log.Info("Getting the order by table within system")
	restaurant := c.Locals("restaurant").(*rest.Restaurant)
	order := squareup.GetOrderByTable(restaurant.LocationID, tableID)
	query := primitive.M{"square_up_order_id": order.SquareUpOrderId}
	repository.UpdateBy(query, order)
	return &dto.CreateOrderRes{
		Data: order,
	}
}

// mapDTOToModel maps the CreateOrderReq DTO to the Order model
func mapDTOToModel(req dto.CreateOrderReq) models.Order {

	order := models.Order{
		OpenedAt: req.OpenedAt,
		IsClosed: req.IsClosed,
		Table:    req.Table,
	}

	for _, item := range req.Items {
		orderItem := models.OrderItem{
			Name:      item.Name,
			Comment:   item.Comment,
			UnitPrice: item.UnitPrice,
			Quantity:  strconv.Itoa(item.Quantity),
			Amount:    item.Amount,
		}

		for _, discount := range item.Discounts {
			orderItem.Discounts = append(orderItem.Discounts, models.Discount{
				Name:         discount.Name,
				IsPercentage: discount.IsPercentage,
				Value:        discount.Value,
				Amount:       discount.Amount,
			})
		}

		for _, modifier := range item.Modifiers {
			orderItem.Modifiers = append(orderItem.Modifiers, models.Modifier{
				Name:      modifier.Name,
				UnitPrice: modifier.UnitPrice,
				Quantity:  strconv.Itoa(modifier.Quantity),
				Amount:    modifier.Amount,
			})
		}

		order.Items = append(order.Items, orderItem)
	}

	order.Totals = models.OrderTotals{
		Discounts:     req.Totals.Discounts,
		Due:           req.Totals.Due,
		Tax:           req.Totals.Tax,
		ServiceCharge: req.Totals.ServiceCharge,
		Paid:          req.Totals.Paid,
		Tips:          req.Totals.Tips,
		Total:         req.Totals.Total,
	}

	return order
}
