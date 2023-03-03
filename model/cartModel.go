package model


// Cart represents a user's shopping cart
type Cart struct {
	ID          string  `json:"id" bson:"_id,omitempty"`
	UserID      string  `json:"userId" bson:"userId,omitempty"`
	InventoryID string  `json:"inventoryId" bson:"inventoryId,omitempty"`
	Quantity    int     `json:"quantity" bson:"quantity"`
	OrderTotal  float64 `json:"orderTotal" bson:"orderTotal"`
	SellerID    string  `json:"sellerId" bson:"sellerId,omitempty"`
	BuyerID     string  `json:"buyerId" bson:"buyerId,omitempty"`
}

