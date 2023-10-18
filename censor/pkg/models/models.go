package models

type BannedWord struct {
	ID   int    `json:"ID,omitempty"`
	Word string `json:"word,omitempty"`
}
