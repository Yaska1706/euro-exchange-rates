package api

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	db     *sql.DB
	router *mux.Router
}

func NewServer(router *mux.Router) *server {
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
