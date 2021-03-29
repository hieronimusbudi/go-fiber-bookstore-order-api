package services

import (
	"time"

	resterrors "github.com/hieronimusbudi/go-bookstore-utils/rest_errors"
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/domain/items"
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/domain/orders"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ordersService struct{}
type ordersServiceInterface interface {
	Create(orders.Order) (*orders.Order, resterrors.RestErr)
}

var (
	OrdersService ordersServiceInterface = &ordersService{}
)

func (s *ordersService) Create(order orders.Order) (*orders.Order, resterrors.RestErr) {
	order.ExpiresAt = time.Now().UTC().Format("2006-01-02 15:04:05")

	if order.ItemId == "" {
		validateErr := resterrors.NewBadRequestError("no item id")
		return nil, validateErr
	}

	itemID, idErr := primitive.ObjectIDFromHex(order.ItemId)
	if idErr != nil {
		return nil, resterrors.NewInternalServerError("create object id error", idErr)
	}

	item := items.Item{ID: itemID}
	itemErr := item.GetSingle()
	if itemErr != nil {
		return nil, itemErr
	}

	if err := order.Save(); err != nil {
		return nil, err
	}

	return &order, nil
}
