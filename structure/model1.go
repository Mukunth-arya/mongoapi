package models

import (
	"time"

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

var data1 = []*data{
	&data{
		NAME:        "DARKMELOWE",
		COCOA:       70,
		CALORIE:     6,
		INGREDIENTS: "cocoa solids, cocoa butter, sugar, milk solids",
		MFD:         time.Now().UTC().Day(),
		EXPDT:       20 / 01 / 2025,
	},
	&data{
		ID:          0002,
		NAME:        "DARKNUTEE",
		COCOA:       70,
		CALORIE:     6,
		INGREDIENTS: "cocoa solids, cocoa butter, sugar, milk solids",
		MFD:         time.Now().UTC().Day(),
		EXPDT:       20 / 01 / 2025,
	},
}
