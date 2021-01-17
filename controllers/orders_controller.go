package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	resterrors "github.com/hieronimusbudi/go-bookstore-utils/rest_errors"
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/domain/orders"
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/services"
)

func CreateOrder(c *fiber.Ctx) error {
	order := new(orders.Order)
	if err := c.BodyParser(order); err != nil {
		restErr := resterrors.NewBadRequestError("invalid json body")
		c.Status(restErr.Status()).JSON(restErr)
		return nil
	}

	result, saveErr := services.OrdersService.CreateOrder(*order)
	if saveErr != nil {
		c.Status(saveErr.Status()).JSON(saveErr)
		return nil
	}

	c.Status(http.StatusCreated).JSON(result)
	return nil
}

// func CreateItem(c *gin.Context) {
// 	var item items.Item
// 	if err := c.ShouldBindJSON(&item); err != nil {
// 		restErr := resterrors.NewBadRequestError("invalid json body")
// 		c.JSON(restErr.Status(), restErr)
// 		return
// 	}

// 	result, saveErr := services.ItemsService.Create(item)
// 	if saveErr != nil {
// 		c.JSON(saveErr.Status(), saveErr)
// 		return
// 	}

// 	c.JSON(http.StatusCreated, result)
// }
