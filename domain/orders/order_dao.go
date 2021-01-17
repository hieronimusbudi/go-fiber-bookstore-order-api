package orders

import (
	"context"
	"errors"

	resterrors "github.com/hieronimusbudi/go-bookstore-utils/rest_errors"
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/datasources/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (order *Order) Save() resterrors.RestErr {
	insertResult, insertErr := mongodb.OrdersCollection.InsertOne(context.TODO(), order)
	if insertErr != nil {
		return resterrors.NewInternalServerError("error when trying to save order", insertErr)
	}

	orderId, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return resterrors.NewInternalServerError("error when trying to save order", errors.New("error when trying to save order"))
	}
	order.ID = orderId.Hex()
	return nil
}
