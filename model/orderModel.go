package model
import (
	"time"
)

type Order struct {
	ID            string    `bson:"_id,omitempty"`
	UserID        string    `bson:"userId,omitempty"`
	InventoryID   string    `bson:"inventoryId,omitempty"`
	BuyerID       string    `bson:"buyerId,omitempty"`
	SellerID      string    `bson:"sellerId,omitempty"`
	Quantity       int      `bson:"quantity"`
	TotalPrice    float64   `bson:"totalPrice"`
	PaymentMethod  BankDetails    `bson:"paymentMethod,omitempty"`
	OrderDate     time.Time `bson:"orderDate"`
	DeliveryDate  time.Time `bson:"deliveryDate,omitempty"`
	Address        Address  `bson:"address,omitempty"`
	CreatedAt     time.Time `bson:"createdAt"`
	UpdatedAt     time.Time `bson:"updatedAt"`
}
