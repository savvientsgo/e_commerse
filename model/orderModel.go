package model
import (
	"time"
)

type Order struct {
	ID            string          `json:"_id" bson:"_id,omitempty"`
	UserID        string          `json:"userId" bson:"userId,omitempty"`
	InventoryID   string          `json:"inventoryId" bson:"inventoryId,omitempty"`
	BuyerID       string          `json:"buyerId" bson:"buyerId,omitempty"`
	SellerID      string          `json:"sellerId" bson:"sellerId,omitempty"`
	Quantity       int            `json:"quantity" bson:"quantity"`
	TotalPrice    float64         `json:"totalPrice" bson:"totalPrice"`
	PaymentMethod  BankDetails    `json:"paymentMethod" bson:"paymentMethod,omitempty"`
	OrderDate     time.Time       `json:"orderDate" bson:"orderDate"`
	DeliveryDate  time.Time       `json:"deliveryDate" bson:"deliveryDate,omitempty"`
	Address        Address        `json:"address" bson:"address,omitempty"`
	CreatedAt     time.Time       `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt" bson:"updatedAt"`
}
