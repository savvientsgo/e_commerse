package model

type Cart struct {
	ID          string  `bson:"_id,omitempty"`
	UserID      string  `bson:"userId,omitempty"`
	InventoryID string  `bson:"inventoryId,omitempty"`
	Quantity    uint    `bson:"quantity"`
	OrderTotal  float64 `bson:"orderTotal"`
	SellerID    string  `bson:"sellerId,omitempty"`
	BuyerID     string  `bson:"buyerId,omitempty"`
}
