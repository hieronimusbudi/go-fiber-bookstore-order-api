package items

import (
	"context"
	"errors"
	"fmt"
	"strings"

	resterrors "github.com/hieronimusbudi/go-bookstore-utils/rest_errors"
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/datasources/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (item *Item) Save() resterrors.RestErr {
	insertResult, insertErr := mongodb.ItemsCollection.InsertOne(context.TODO(), item)
	if insertErr != nil {
		return resterrors.NewInternalServerError("error when trying to save item", insertErr)
	}

	itemID, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return resterrors.NewInternalServerError("error when trying to save item", errors.New("error when trying to save item"))
	}
	item.ID = itemID.Hex()
	return nil
}

func (item *Item) GetSingle() resterrors.RestErr {
	// itemID, idErr := primitive.ObjectIDFromHex(item.ID)
	// if idErr != nil {
	// 	return resterrors.NewInternalServerError("create object id error", idErr)
	// }

	filter := bson.M{"sourceid": item.ID}
	findErr := mongodb.ItemsCollection.FindOne(context.TODO(), filter).Decode(&item)
	if findErr != nil {
		if strings.Contains(findErr.Error(), "404") {
			return resterrors.NewNotFoundError(fmt.Sprintf("no item found with id %s", item.ID))
		}
		return resterrors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", item.ID), errors.New(findErr.Error()))
	}

	return nil
}
