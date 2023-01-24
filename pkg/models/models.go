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