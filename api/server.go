package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"
)

type server struct {
	db     *sql.DB
	router *http.ServeMux
}

func NewServer(router *http.ServeMux) *server {
	return &server{
		router: router,
	}
}

func (s *server) Run() {

	serve := &http.Server{
		Addr:              "127.0.0.1:5000",
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Handler:           s.routes(),
	}
	serve.ListenAndServe()
	log.Print("App Running")

}
