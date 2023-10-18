package main

import (
	"comments/pkg/api"
	"comments/pkg/middleware"
	"comments/pkg/storage"
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

	port := goDotEnvVariable("PORT_COMMENT")
	connstr := goDotEnvVariable("CONNSTR")

	db, err := storage.New(connstr)

	if err != nil {
		log.Fatal(err)
	}

	srv.db = *db
	srv.api = api.New(srv.db)

	srv.api.Router().Use(middleware.Middle)

	log.Print("start comments server http://127.0.0.1" + port)

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
