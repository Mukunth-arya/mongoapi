package controller

import (
	"context"
	"fmt"
	"log"

	"name/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Connectionvalue = "mongodb+srv://mukunth:mukunth@mycluster.jptcn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbname = "smokedetector"
const colname = "smokecol"

var data *mongo.Collection

func init() {

	c1 := options.Client().ApplyURI(Connectionvalue)
	// connect to database:
	c2, err := mongo.Connect(context.TODO(), c1)
	if err != nil {

		log.Fatal(err)
	}
	fmt.Println("mongodb connection is superb")

	c3 := c2.Database(dbname).Collection(colname)

	fmt.Println("All are ready", c3)
}

func insertdataone(datavalue models.Data) {
	cur, err := data.InsertOne(context.Background(), datavalue)
	if err != nil {

		log.Fatal(err)
	}
	fmt.Println("one data were inserted", cur.InsertedID)

}
func efficentparameter(valueID string) {

	cur1, _ := primitive.ObjectIDFromHex(valueID)
	filter := bson.M{"_id": cur1}
	update := bson.M{"$set": bson.M{"Trustworthy": true}}
	cur, err := data.UpdateOne(context.Background(), filter, update)
	if err != nil {

		log.Fatal(err)
	}
	fmt.Println("were all the data is modified", cur.ModifiedCount)

}
func deleteone(valueID string) {

	cur1, _ := primitive.ObjectIDFromHex(valueID)
	filter := bson.M{"_id": cur1}
	cur, err := data.DeleteOne(context.Background(), filter)
	if err != nil {

		log.Fatal(err)
	}
	fmt.Println("here comes the deletecount", cur)

}

func deleteall() int64 {

	cur, err := data.DeleteMany(context.Background(), bson.D{})
	if err != nil {

		log.Fatal(err)
	}
	fmt.Println("here comes the Deletedcount", cur.DeletedCount)
	return cur.DeletedCount

}
func getallsource() {

	cur, err := data.Find(context.Background(), bson.M{})
	if err != nil {

		log.Fatal(err)
	}
	var datahold []primitive.M
	for cur.Next(context.Background()) {

		var files bson.M
		err := cur.Decode(&files)
		if err != nil {

			log.Fatal(err)
		}
		datahold = append(datahold, files)
	}
	defer cur.Close(context.Background())
}
