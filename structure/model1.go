package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type data struct {
	ID          primitive.ObjectID `json: "_id,omitempty" bson:"id_,omitempty`
	NAME        string             `json: "name,omitempty"`
	COCOA       int                `json: "cocoa,omitempty"`
	CALORIE     int                `json: "calorie,omitempty"`
	INGREDIENTS string             `json: "ingredients,omitempty"`
	MFD         int                `json: "-"`
	EXPDT       int                `json: "-"`
}
