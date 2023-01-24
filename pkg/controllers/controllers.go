package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"search_engine_task/pkg"
	"search_engine_task/pkg/dbconn"
	"search_engine_task/pkg/models"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type PagesByScore []models.Result

func (u PagesByScore) Len() int {
	return len(u)
}

func (u PagesByScore) Swap(i, j int) {

	u[i], u[j] = u[j], u[i]

}

func (u PagesByScore) Less(i, j int) bool {

	return u[i].Score < u[j].Score

}

type Interface interface {
	Len() int

	Less(i, j int) bool

	Swap(i, j int)
}

type Service interface {
	Get()
	Insert()
	Check()
	Routes()
}

type Controller struct {
	service Service
	// db *dbconn.DB
}

func (ct *Controller) Get(c *gin.Context) {

	var words models.Keywords

	if err := c.BindJSON(&words); err != nil {

		return

	}

	temp := pkg.GetAllCollection()

	pages := []models.Page{}

	for _, p := range temp {

		var s models.Page

		bsonBytes, _ := bson.Marshal(p)

		bson.Unmarshal(bsonBytes, &s)

		pages = append(pages, s)

	}

	res := []models.Result{}

	for i := 0; i < len(pages); i++ {
		var tempRs models.Result

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

	for i := len(ans) - 1; i >= 0; i-- {

		fmt.Println(ans[i])

	}

}

func (ct *Controller) Insert(c *gin.Context) {

	var newPage models.Page

	if err := c.BindJSON(&newPage); err != nil {

		return

	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, insertErr := dbconn.Dbconn().InsertOne(ctx, newPage)

	if insertErr != nil {

		println("InsertONE Error:", insertErr)

		os.Exit(1)

	}

	c.IndentedJSON(http.StatusCreated, newPage)

}

func (ct *Controller) Check(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, "Alive")
}

func (ct *Controller) Routes(router *gin.RouterGroup) {
	router.GET("/", ct.Check)
	router.POST("insert", ct.Insert)
	router.GET("get", ct.Get)
}
