package api

import (
	"database/sql"
	"log"
	"net/http"
	"os"
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

	ADDR := os.Getenv("LISTEN_ADDRESS")
	PORT := os.Getenv("LISTEN_PORT")
	ListenADDR := ADDR + ":" + PORT

	serve := &http.Server{
		Addr:              ListenADDR,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Handler:           s.routes(),
	}
	log.Print("App Running")
	serve.ListenAndServe()

}
