package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type page struct { //API 1
	Title    string   `json:"title"`
	Keywords []string `json:"keywords"`
}

type user_keywords struct { // API 2
	User_keywords []string `json:"user_keywords"`
}

// var res page

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

}

func main() {
	router := gin.Default()
	router.POST("/api1", savePage)
	router.POST("/api2", getResult)

	router.Run("localhost:8080")
}
