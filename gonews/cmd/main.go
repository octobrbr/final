package main

import (
	"gonews/pkg/api"
	"gonews/pkg/middleware"
	"gonews/pkg/models"
	"gonews/pkg/rss"
	"gonews/pkg/storage"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type server struct {
	db  storage.DB
	api *api.API
}

func main() {
	var srv server

	port := goDotEnvVariable("PORT_NEWS")
	connstr := goDotEnvVariable("CONNSTR")

	db, err := storage.New(connstr)

	if err != nil {
		log.Fatal(err)
	}

	srv.db = *db
	srv.api = api.New(srv.db)

	chanPosts := make(chan []models.Post)
	chanErrs := make(chan error)

	go func() {
		err := rss.GetNews(chanPosts, chanErrs)
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		for posts := range chanPosts {
			if err := srv.db.AddPosts(posts); err != nil {
				chanErrs <- err
			}
		}
	}()

	go func() {
		for err := range chanErrs {
			log.Println(err)
		}
	}()

	srv.api.Router().Use(middleware.Middle)

	log.Print("start news server http://127.0.0.1" + port)

	err = http.ListenAndServe(port, srv.api.Router())
	if err != nil {
		log.Fatal("Couldnt start server. Error:", err)
	}
}

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
