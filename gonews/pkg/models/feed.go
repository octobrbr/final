package models

import "encoding/xml"

// type Feed struct {
// 	XMLName xml.Name `xml:"rss"`
// 	Chanel  Channel  `xml:"channel"`
// }

// type Channel struct {
// 	Title       string    `xml:"title"`
// 	Description string    `xml:"description"`
// 	Link        string    `xml:"link"`
// 	Items       []ItemRSS `xml:"item"`
// }

// type ItemRSS struct {
// 	Title       string `xml:"title"`
// 	Description string `xml:"description"`
// 	PubDate     string `xml:"pubDate"`
// 	Link        string `xml:"link"`
// }

type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Chanel  Channel  `xml:"channel"`
}

// структура отдельной новости из RSS
type Channel struct {
	Title       string    `xml:"title"`
	Description string    `xml:"description"`
	Link        string    `xml:"link"`
	Items       []ItemRSS `xml:"item"`
}

// структура отдельной новости из RSS
type ItemRSS struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Link        string `xml:"link"`
}

