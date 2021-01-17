package services

import (
	"time"

	resterrors "github.com/hieronimusbudi/go-bookstore-utils/rest_errors"
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/domain/items"
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/domain/orders"
)

type ordersService struct{}
type ordersServiceInterface interface {
	CreateOrder(orders.Order) (*orders.Order, resterrors.RestErr)
}

var (
	OrdersService ordersServiceInterface = &ordersService{}
)

func (s *ordersService) CreateOrder(order orders.Order) (*orders.Order, resterrors.RestErr) {
	order.ExpiresAt = time.Now().UTC().Format("2006-01-02 15:04:05")

	if order.ItemId == "" {
		validateErr := resterrors.NewBadRequestError("no item id")
		return nil, validateErr
	}

	item := items.Item{ID: order.ItemId}
	itemErr := item.GetSingle()
	if itemErr != nil {
		return nil, itemErr
	}

	if err := order.Save(); err != nil {
		return nil, err
	}

	return &order, nil
}
