package controllers

import (
	"context"
	"net/http"
	"os"
	"search_engine_task/pkg"
	"search_engine_task/pkg/dbconn"
	"time"

	"github.com/gin-gonic/gin"
)

type PageController struct {
}

func (p PageController) SavePage(c *gin.Context) {

	var newPage pkg.Page

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
