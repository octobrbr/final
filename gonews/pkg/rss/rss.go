package rss

import (
	"encoding/json"
	"encoding/xml"
	"gonews/pkg/models"
	"log"
	"net/http"
	"os"
	"time"

	strip "github.com/grokify/html-strip-tags-go"
)

type Item struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	Content string `xml:"description"`
	PubDate string `xml:"pubDate"`
}

type Channel struct {
	Items []Item `xml:"channel>item"`
}

type config struct {
	Rss           []string `json:"rss"`
	RequestPeriod int      `json:"request_period"`
}

func GetNews(chanPosts chan<- []models.Post, chanErrs chan<- error) error {
	file, err := os.Open("./cmd/config.json")
	if err != nil {
		return err
	}

	var conf config

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		return err
	}

	log.Println("start read rss-feed")
	for i, r := range conf.Rss {
		go func(r string, i int, chanPosts chan<- []models.Post, chanErrs chan<- error) {
			for {
				log.Println("start  goroutine", i, "url: ", r)
				p, err := GetRss(r)
				if err != nil {
					chanErrs <- err
					time.Sleep(time.Second * 10)
					continue
				}
				chanPosts <- p
				log.Println("insert posts from goroutine", i, "url: ", r)
				log.Println("Goroutine ", i, ": waiting next iteration")
				time.Sleep(time.Duration(conf.RequestPeriod) * time.Second * 15)
			}
		}(r, i, chanPosts, chanErrs)
	}
	return nil
}

func GetRss(url string) ([]models.Post, error) {
	var c Channel

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	err = xml.NewDecoder(res.Body).Decode(&c)
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	for _, i := range c.Items {
		var p models.Post
		p.Title = i.Title
		p.Content = i.Content
		p.Content = strip.StripTags(p.Content)
		p.Link = i.Link

		t, err := time.Parse(time.RFC1123, i.PubDate)
		if err != nil {
			t, err = time.Parse(time.RFC1123Z, i.PubDate)
		}
		if err != nil {
			t, err = time.Parse("Mon, _2 Jan 2006 15:04:05 -0700", i.PubDate)
		}
		if err == nil {
			p.PubTime = t.Unix()
		}

		posts = append(posts, p)
	}
	return posts, nil
}
