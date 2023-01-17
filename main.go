package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var collection *mongo.Collection

type page struct { //API 1
	Title    string   `json:"title"`
	Keywords []string `json:"keywords"`
}

type keywords struct {
	User_keywords []string
}

func savePage(c *gin.Context) {
	var newPage page
	if err := c.BindJSON(&newPage); err != nil {
		return
	}
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := client.Database("searchengine").Collection("pages")
	_, insertErr := col.InsertOne(ctx, newPage)
	if insertErr != nil {
		fmt.Println("InsertONE Error:", insertErr)
		os.Exit(1)
	}
	c.IndentedJSON(http.StatusCreated, newPage)
}

func getResult(c *gin.Context) {
	var words keywords
	if err := c.BindJSON(&words); err != nil {
		return
	}
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}

	col := client.Database("searchengine").Collection("pages")
	cursor, err := col.Find(context.TODO(), bson.D{})
	
	


}


// func displayWords(c *gin.Context) {
// 	var words keywords
// 	if err := c.BindJSON(&words); err != nil {
// 		return
// 	}
// 	c.IndentedJSON(http.StatusCreated, words)
// }







func main() {
	router := gin.Default()
	router.POST("/api1", savePage)
	router.POST("/api2", getResult)
	// router.POST("/api3", displayWords)
	router.Run("localhost:8080")
}