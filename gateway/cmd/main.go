package main

import (
	"final/gateway/pkg/api"
	"final/gateway/pkg/middleware"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type server struct {
	api *api.API
}

func main() {
	var srv server

	portGateway := goDotEnvVariable("PORT_GATEWAY")
	portNews := goDotEnvVariable("PORT_NEWS")
	portComment := goDotEnvVariable("PORT_COMMENT")
	portCensor := goDotEnvVariable("PORT_CENSOR")

	srv.api = api.New(portNews, portCensor, portComment)

	srv.api.Router().Use(middleware.Middle)

	log.Print("start gateway server http://127.0.0.1" + portGateway + "/news")

	err := http.ListenAndServe(portGateway, srv.api.Router())
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
