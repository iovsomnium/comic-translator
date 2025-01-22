package api

import (
	"github.com/gorilla/mux"
	"github.com/iovsomnium/comic-translator/src/config"
	"github.com/iovsomnium/comic-translator/src/service"
)

func SetupRouter(cfg *config.Config) *mux.Router {
    router := mux.NewRouter()

    // 만화책 API
    router.HandleFunc("/comics/{id}/translate", service.TranslateComic).Methods("POST")

    return router
}
