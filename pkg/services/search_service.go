package services

import (
	"context"
	"search_engine_task/pkg/external"
	"search_engine_task/pkg/models"
	"sort"
	"strings"
)

type SearchService struct {
	dbClient external.IDBClient
}

func NewSearchService(dbParam external.IDBClient) *SearchService {
	return &SearchService{dbClient: dbParam}
}

func (s SearchService) SavePage(ctx context.Context, page models.Page) error {
	return s.dbClient.InsertOnePage(ctx, page)
}

func (s SearchService) GetResult(ctx context.Context, words models.Keywords) ([]string, error) {
	pages := s.dbClient.GetAllCollection()

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

	sort.Stable(models.PagesByScore(res))

	ans := []string{}

	for i := 0; i < len(res); i++ {

		ans = append(ans, res[i].Title)

	}

	return ans, nil
}
