package model

import "time"

type Inventory struct {
	ID          string    `bson:"_id,omitempty"`
	Name        string    `bson:"name"`
	Type        string    `bson:"type"`
	Description string    `bson:"description,omitempty"`
	Price       float64   `bson:"price"`
	Quantity    uint      `bson:"quantity"`
	SellerID    string    `bson:"sellerId,omitempty"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
	Condition   string    `bson:"condition,omitempty"`
}
