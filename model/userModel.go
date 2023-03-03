package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	FirstName    string             `json:"firstName" bson:"firstName"`
	LastName     string             `json:"lastName" bson:"lastName"`
	Email        string             `json:"email" bson:"email"`
	UserName     string             `json:"userName" bson:"userName"`
	PasswordHash string             `json:"passwordHash" bson:"passwordHash"`
	Address      Address            `json:"address" bson:"address"`
	BankDetail   BankDetail         `json:"bankDetail" bson:"bankDetail"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
}

//  Address and BankDetails

type Address struct {
	Street     string `json:"street" bson:"street"`
	City       string `json:"city" bson:"city"`
	State      string `json:"state" bson:"state"`
	Country    string `json:"country" bson:"country"`
	PostalCode string `json:"postalCode" bson:"postalCode"`
}

type BankDetail struct {
	BankName    string `json:"bankName" bson:"bankName"`
	AccountName string `json:"accountName" bson:"accountName"`
	AccountNo   string `json:"accountNo" bson:"accountNo"`
	RoutingNo   string `json:"routingNo" bson:"routingNo"`
}
