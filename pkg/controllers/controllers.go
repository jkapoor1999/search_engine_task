package controllers

import (
	"context"
	"net/http"
	"search_engine_task/pkg/models"
	"search_engine_task/pkg/services"
	"time"

	"github.com/gin-gonic/gin"
)

type Interface interface {
	Len() int

	Less(i, j int) bool

	Swap(i, j int)
}

type IController interface {
	Get()
	Insert()
	Check()
	Routes()
}

type Controller struct {
	searchService services.ISearchService
}

func NewController(searchService services.ISearchService) *Controller {
	return &Controller{searchService: searchService}
}

func (ct *Controller) Get(c *gin.Context) {

	var words models.Keywords

	if err := c.BindJSON(&words); err != nil {

		return

	}

	ans, err := ct.searchService.GetResult(c, words)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, ans)
}

func (ct *Controller) Insert(c *gin.Context) {

	var newPage models.Page

	if err := c.BindJSON(&newPage); err != nil {

		return

	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	insertErr := ct.searchService.SavePage(ctx, newPage)

	if insertErr != nil {

		c.IndentedJSON(http.StatusInternalServerError, insertErr)
		return

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
	// ctx.IndentedJSON(http.StatusCreated)
}
