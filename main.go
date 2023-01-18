package main

import (
	"context"

	"sort"

	"strings"

	"fmt"

	"net/http"

	"os"

	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type page struct { //API 1

	Title string `json:"title"`

	Keywords []string `json:"keywords"`
}

type keywords struct {
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

func getAllCollection() []bson.M {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {

		fmt.Println("Mongo.connect() ERROR: ", err)

		os.Exit(1)

	}

	col := client.Database("searchengine").Collection("pages")

	cur, err := col.Find(context.Background(), bson.D{})

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

func getResult(c *gin.Context) {

	var words keywords

	if err := c.BindJSON(&words); err != nil {

		return

	}

	temp := getAllCollection()

	pages := []page{}

	for _, p := range temp {

		var s page

		bsonBytes, _ := bson.Marshal(p)

		bson.Unmarshal(bsonBytes, &s)

		pages = append(pages, s)

	}

	// for _, p := range pages {

	//  println(p.Title)

	// }

	res := []Result{}

	for i := 0; i < len(pages); i++ {
                var tempRs Result

                tempScore := 0
		for j := 0; j < len(pages[i].Keywords); j++ {

			

			for k := 0; k < len(words.User_keywords); k++ {

				if strings.EqualFold(pages[i].Keywords[j], words.User_keywords[k]) {

					tempScore += (10 - k) * (10 - j)

				}

			}

			

		}
        if tempScore > 0 {

            tempRs.Title = pages[i].Title

            tempRs.Score = tempScore

        }

        res = append(res, tempRs)

	}

	sort.Stable(PagesByScore(res))

	ans := []string{}

	for i := 0; i < len(res); i++ {

		ans = append(ans, res[i].Title)

	}

	for i := len(ans)-1; i >= 0; i-- {

		fmt.Println(ans[i])

	}

}

func main() {

	router := gin.Default()

	router.POST("/api1", savePage)

	router.POST("/api2", getResult)

	router.Run("localhost:8084")

}
