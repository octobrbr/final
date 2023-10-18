package api

import (
	"censor/pkg/storage"
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

type API struct {
	r  *mux.Router
	db storage.DB
}

func New(db storage.DB) *API {
	api := API{
		r:  mux.NewRouter(),
		db: db,
	}
	api.r = mux.NewRouter()
	api.endpoints()
	return &api
}

func (api *API) Router() *mux.Router {
	return api.r
}

func (api *API) endpoints() {
	api.r.HandleFunc("/comments/check", api.addCommentHandler).Methods(http.MethodPost, http.MethodOptions)
}

func (api *API) addCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	text := struct {
		Content string
	}{}
	err := json.NewDecoder(r.Body).Decode(&text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	blacklist, err := api.db.GetBlackList()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, banWord := range blacklist {
		matched, err := regexp.MatchString(banWord.Word, text.Content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if matched {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
