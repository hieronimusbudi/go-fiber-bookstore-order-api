package orders

import (
	"context"
	"errors"
	"fmt"
	"strings"

	resterrors "github.com/hieronimusbudi/go-bookstore-utils/rest_errors"
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/datasources/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (o *Order) Save() resterrors.RestErr {
	insertResult, insertErr := mongodb.OrdersCollection.InsertOne(context.TODO(), o)
	if insertErr != nil {
		return resterrors.NewInternalServerError("error when trying to save order", insertErr)
	}

	orderId, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return resterrors.NewInternalServerError("error when trying to save order", errors.New("error when trying to save order"))
	}
	o.ID = orderId
	return nil
}

func (o *Order) GetAll() ([]Order, resterrors.RestErr) {
	filter := bson.D{{}}
	cur, findErr := mongodb.OrdersCollection.Find(context.TODO(), filter)
	if findErr != nil {
		return nil, resterrors.NewInternalServerError(findErr.Error(), errors.New(findErr.Error()))
	}

	orders := []Order{}
	//Map result to slice
	for cur.Next(context.TODO()) {
		t := Order{}
		err := cur.Decode(&t)
		if err != nil {
			return nil, resterrors.NewInternalServerError(err.Error(), errors.New(err.Error()))
		}
		orders = append(orders, t)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(orders) == 0 {
		return nil, resterrors.NewInternalServerError(mongo.ErrNoDocuments.Error(), mongo.ErrNoDocuments)
	}

	return orders, nil
}

func (o *Order) GetSingle() resterrors.RestErr {
	filter := bson.M{"_id": o.ID}
	findErr := mongodb.OrdersCollection.FindOne(context.TODO(), filter).Decode(&o)
	if findErr != nil {
		if strings.Contains(findErr.Error(), "404") {
			return resterrors.NewNotFoundError(fmt.Sprintf("no order found with id %s", o.ID))
		}
		return resterrors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", o.ID), errors.New(findErr.Error()))
	}

	return nil
}

func (o *Order) Update() resterrors.RestErr {
	filter := bson.M{"_id": o.ID}
	oByte, err := bson.Marshal(*o)
	if err != nil {
		return resterrors.NewInternalServerError("error when trying to marshal item", err)
	}

	var update bson.M
	err = bson.Unmarshal(oByte, &update)
	if err != nil {
		return resterrors.NewInternalServerError("error when trying to unmarshal item", err)
	}

	_, updateErr := mongodb.OrdersCollection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: update}})
	if updateErr != nil {
		return resterrors.NewInternalServerError("error when trying to update order", updateErr)
	}

	return nil
}

func (o *Order) Delete() resterrors.RestErr {
	filter := bson.M{"_id": o.ID}

	_, deleteErr := mongodb.ItemsCollection.DeleteOne(context.TODO(), filter)
	if deleteErr != nil {
		return resterrors.NewInternalServerError("error when trying to delete order", deleteErr)
	}

	return nil
}
