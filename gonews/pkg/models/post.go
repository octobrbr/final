package models

// структура отдельной новости для БД
type Post struct {
	ID      int    `json:"ID,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	PubTime int64  `json:"pubTime,omitempty"`
	Link    string `json:"link,omitempty"`
}

type Pagination struct {
	NumOfPages int `json:"numOfPages,omitempty"`
	Page       int `json:"page,omitempty"`
	Limit      int `json:"limit,omitempty"`
}
