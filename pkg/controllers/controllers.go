package controllers

import (
	"fmt"
	"search_engine_task/pkg"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetResult(c *gin.Context) {

	var words pkg.Keywords

	if err := c.BindJSON(&words); err != nil {

		return

	}

	temp := pkg.GetAllCollection()

	pages := []pkg.Page{}

	for _, p := range temp {

		var s pkg.Page

		bsonBytes, _ := bson.Marshal(p)

		bson.Unmarshal(bsonBytes, &s)

		pages = append(pages, s)

	}

	res := []pkg.Result{}

	for i := 0; i < len(pages); i++ {
		var tempRs pkg.Result

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

	sort.Stable(pkg.PagesByScore(res))

	ans := []string{}

	for i := 0; i < len(res); i++ {

		ans = append(ans, res[i].Title)

	}

	for i := len(ans) - 1; i >= 0; i-- {

		fmt.Println(ans[i])

	}

}
