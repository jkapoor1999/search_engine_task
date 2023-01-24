package models

type Page struct { 

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