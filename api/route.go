package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) routes() *mux.Router {
	router := s.router

	router.HandleFunc("/status", s.ApiStatus())
	router.HandleFunc("/rates/latest", s.GetLatest()).Methods("GET")
	router.HandleFunc("/rates/analyze", s.AnalyzeRates())
	// router.HandleFunc("/rates/{date}", s.GetSpecificDate())
	return router
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}
