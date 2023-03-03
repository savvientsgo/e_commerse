package model

import "time"

type Inventory struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Type        string    `json:"type" bson:"type"`
	Description string    `json:"description" bson:"description,omitempty"`
	Price       float64   `json:"price" bson:"price"`
	Quantity     int      `json:"quantity" bson:"quantity"`
	SellerID    string    `json:"sellerId" bson:"sellerId,omitempty"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
	Condition   string    `json:"condition" bson:"condition,omitempty"`
}
