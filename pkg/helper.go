package pkg

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Page struct { //API 1

	Title string `json:"title"`

	Keywords []string `json:"keywords"`
}

type Keywords struct {
	User_keywords []string `json:"user_keywords"`
}

type Result struct {
	Title string

	Score int
}

type Interface interface {
	Len() int

	Less(i, j int) bool

	Swap(i, j int)
}

type PagesByScore []Result

func (u PagesByScore) Len() int {
	return len(u)
}

func (u PagesByScore) Swap(i, j int) {

	u[i], u[j] = u[j], u[i]

}

func (u PagesByScore) Less(i, j int) bool {

	return u[i].Score < u[j].Score

}

func Dbconn() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {

		fmt.Println("Mongo.connect() ERROR: ", err)

		os.Exit(1)

	}

	return client.Database("searchengine").Collection("pages")
}

func GetAllCollection() []bson.M {

	cur, err := Dbconn().Find(context.Background(), bson.D{})

	if err != nil {

		os.Exit(1)

	}

	var temp []bson.M

	for cur.Next(context.Background()) {

		var p bson.M

		err := cur.Decode(&p)

		if err != nil {

			os.Exit(1)

		}

		temp = append(temp, p)

	}

	defer cur.Close(context.Background())

	return temp

}
