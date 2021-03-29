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
	"go.mongodb.org/mongo-driver/mongo"
)

func (i *Item) Save() resterrors.RestErr {
	insertResult, insertErr := mongodb.ItemsCollection.InsertOne(context.TODO(), i)
	if insertErr != nil {
		return resterrors.NewInternalServerError("error when trying to save item", insertErr)
	}

	itemID, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return resterrors.NewInternalServerError("error when trying to save item", errors.New("error when trying to save item"))
	}
	i.ID = itemID
	return nil
}

func (i *Item) GetAll() ([]Item, resterrors.RestErr) {
	filter := bson.D{{}}
	cur, findErr := mongodb.ItemsCollection.Find(context.TODO(), filter)
	if findErr != nil {
		return nil, resterrors.NewInternalServerError(findErr.Error(), errors.New(findErr.Error()))
	}

	items := []Item{}
	//Map result to slice
	for cur.Next(context.TODO()) {
		t := Item{}
		err := cur.Decode(&t)
		if err != nil {
			return nil, resterrors.NewInternalServerError(err.Error(), errors.New(err.Error()))
		}
		items = append(items, t)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(items) == 0 {
		return nil, resterrors.NewInternalServerError(mongo.ErrNoDocuments.Error(), mongo.ErrNoDocuments)
	}

	return items, nil
}

func (i *Item) GetSingle() resterrors.RestErr {
	filter := bson.M{"_id": i.ID}
	findErr := mongodb.ItemsCollection.FindOne(context.TODO(), filter).Decode(&i)
	if findErr != nil {
		if strings.Contains(findErr.Error(), "404") {
			return resterrors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.ID))
		}
		return resterrors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.ID), errors.New(findErr.Error()))
	}

	return nil
}

func (i *Item) Update() resterrors.RestErr {
	filter := bson.M{"_id": i.ID}
	iByte, err := bson.Marshal(*i)
	if err != nil {
		return resterrors.NewInternalServerError("error when trying to marshal item", err)
	}

	var update bson.M
	err = bson.Unmarshal(iByte, &update)
	if err != nil {
		return resterrors.NewInternalServerError("error when trying to unmarshal item", err)
	}

	_, updateErr := mongodb.ItemsCollection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: update}})
	if updateErr != nil {
		return resterrors.NewInternalServerError("error when trying to update item", updateErr)
	}

	return nil
}

func (i *Item) Delete() resterrors.RestErr {
	filter := bson.M{"_id": i.ID}

	_, deleteErr := mongodb.ItemsCollection.DeleteOne(context.TODO(), filter)
	if deleteErr != nil {
		return resterrors.NewInternalServerError("error when trying to delete item", deleteErr)
	}

	return nil
}
