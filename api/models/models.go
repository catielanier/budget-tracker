package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
	Envelopes []Envelope         `json:"envelopes,omitempty" bson:"envelopes,omitempty"`
}

type Envelope struct {
	Name        string  `json:"name,omitempty" bson:"name,omitempty"`
	AllottedAmt float32 `json:"allottedAmt,omitempty" bson:"allottedAmt,omitempty"`
	UsedAmt     float32 `json:"usedAmt,omitempty" bson:"usedAmt,omitempty"`
}

type Transaction struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Recipient string             `json:"recipient,omitempty" bson:"recipient,omitempty"`
	Amount    float32            `json:"amount,omitempty" bson:"amount,omitempty"`
	Envelope  string             `json:"envelope,omitempty" bson:"envelope,omitempty"`
	User      User               `json:"user,omitempty" bson:"user,omitempty"`
}
