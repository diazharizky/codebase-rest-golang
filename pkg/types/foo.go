package types

import "time"

type Foo struct {
	FirstName   string    `json:"firstName" bson:"firstName"`
	LastName    string    `json:"lastName" bson:"lastName"`
	Email       string    `json:"email" bson:"email"`
	PhoneNumber string    `json:"phoneNumber" bson:"phoneNumber"`
	Address     string    `json:"address" bson:"address"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}
