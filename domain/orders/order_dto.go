package orders

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserId    int64              `json:"user_id" bson:"user_id`
	Status    string             `json:"status" bson:"status"`
	ExpiresAt string             `json:"expires_at" bson:"expires_at"`
	ItemId    string             `json:"item_id" bson:"item_id"`
}
