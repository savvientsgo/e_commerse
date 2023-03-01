package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	FirstName    string             `bson:"firstname"`
	LastName     string             `bson:"lastname"`
	Email        string             `bson:"email"`
	UserName     string             `bson:"username"`
	PasswordHash string             `bson:"passwordHash,omitempty"`
	Address      Address            `bson:"address,omitempty"`
	BankDetail   BankDetail         `bson:"bankDetail,omitempty"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
}

//  Address and BankDetails

type Address struct {
	Street     string `bson:"street,omitempty"`
	City       string `bson:"city,omitempty"`
	State      string `bson:"state,omitempty"`
	Country    string `bson:"country,omitempty"`
	PostalCode string `bson:"postalCode,omitempty"`
}

type BankDetail struct {
	BankName    string `bson:"bankName,omitempty"`
	AccountName string `bson:"accountName,omitempty"`
	AccountNo   string `bson:"accountNo,omitempty"`
	RoutingNo   string `bson:"routingNo,omitempty"`
}
