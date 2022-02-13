package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Mukunth-arya/mongoapi/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/monocrypt/options"
)

const Connectionvalue = "mongodb+srv://mukunth:mukunth@mycluster.jptcn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbname = "Trikonechoco"
const colname = "Trikonecell"

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

func (datavalue models.Data) insertdata(w http.ResponseWriter, r *http.Request) {
	insert, err := data.InsertOne(context.Background(), datavalue)
	if err != nil {

		log.Fatal(err)
	}
	fmt.Println("one data were inserted", insert.InsertedID)

}
