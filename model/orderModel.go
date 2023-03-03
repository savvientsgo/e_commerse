package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID        string             `json:"userId" bson:"userId"`
	InventoryID   string             `json:"inventoryId" bson:"inventoryId"`
	BuyerID       string             `json:"buyerId" bson:"buyerId"`
	SellerID      string             `json:"sellerId" bson:"sellerId"`
	Quantity      int                `json:"quantity" bson:"quantity"`
	TotalPrice    float64            `json:"totalPrice" bson:"totalPrice"`
	PaymentMethod BankDetail         `json:"paymentMethod" bson:"paymentMethod"`
	OrderDate     time.Time          `json:"orderDate" bson:"orderDate"`
	DeliveryDate  time.Time          `json:"deliveryDate" bson:"deliveryDate"`
	Address       Address            `json:"address" bson:"address"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt"`
}
