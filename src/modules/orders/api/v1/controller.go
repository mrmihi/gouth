package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"goose/src/global"
	"goose/src/modules/orders/api/v1/dto"
)

func Create(c *fiber.Ctx) error {
	payload := new(dto.CreateOrderReq)
	err := c.BodyParser(payload)
	if err != nil {
		return err
	}
	res := createOrderService(c, *payload)
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.CreateOrderRes]{
		Message: "Orders created successfully",
		Data:    res,
	})
}

func GetById(c *fiber.Ctx) error {
	orderID := c.Params("orderID")
	res := getOrderServiceByID(c, orderID)
	return c.Status(fiber.StatusOK).JSON(global.Response[dto.CreateOrderRes]{
		Message: "Orders retrieved successfully",
		Data:    res,
	})
}

func GetByTable(c *fiber.Ctx) error {
	log.Infof("Getting the order by table within system")
	tableID := c.Params("tableID")
	orders := getOrderServiceByTable(c, tableID)
	return c.JSON(orders)
}
