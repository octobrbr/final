package models

type Comment struct {
	ID      int    `json:"ID,omitempty"`
	NewsID  int    `json:"newsID,omitempty"`
	Content string `json:"content,omitempty"`
	PubTime int64  `json:"pubTime,omitempty"`
}
