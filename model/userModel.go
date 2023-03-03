package model

import (
	"time"
        "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	FirstName    string             `json:"firstName" bson:"firstName"`
	LastName     string             `json:"lastName" bson:"lastName"`
	Email        string             `json:"email" bson:"email"`
	UserName     string             `json:"userName" bson:"userName"`
	PasswordHash string             `json:"passwordHash" bson:"passwordHash,omitempty"`
	Address      Address            `json:"address" bson:"address,omitempty"`
	BankDetail   BankDetail         `json:"bankDetail" bson:"bankDetail,omitempty"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
}

//  Address and BankDetails

type Address struct {
	Street     string `json:"street" bson:"street,omitempty"`
	City       string `json:"city" bson:"city,omitempty"`
	State      string `json:"state" bson:"state,omitempty"`
	Country    string `json:"country" bson:"country,omitempty"`
	PostalCode string `json:"postalCode" bson:"postalCode,omitempty"`
}

type BankDetail struct {
	BankName    string `json:"bankName" bson:"bankName,omitempty"`
	AccountName string `json:"accountName" bson:"accountName,omitempty"`
	AccountNo   string `json:"accountNo" bson:"accountNo,omitempty"`
	RoutingNo   string `json:"routingNo" bson:"routingNo,omitempty"`
}

