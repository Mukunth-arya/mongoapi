package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Data struct {
	ID              primitive.ObjectID `json: "_id,omitempty" bson:"id_,omitempty`
	Productname     string             `json: "productname,omitempty"`
	Countryorgin    string             `json: "countryorgin",omitempty"`
	Batterylife     string             `json: "batterylife,omitempty"`
	protectionglass string             `json: "protectionglass,omitempty"`
	Trustworthy     string             `json: "trustworthy,omitempty"`
}
