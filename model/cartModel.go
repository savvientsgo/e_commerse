package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Cart represents a user's shopping cart
type Cart struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserID      string             `json:"userId" bson:"userId"`
	InventoryID string             `json:"inventoryId" bson:"inventoryId"`
	Quantity    int                `json:"quantity" bson:"quantity"`
	OrderTotal  float64            `json:"orderTotal" bson:"orderTotal"`
	SellerID    string             `json:"sellerId" bson:"sellerId"`
	BuyerID     string             `json:"buyerId" bson:"buyerId"`
}
