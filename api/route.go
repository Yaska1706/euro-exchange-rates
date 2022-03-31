package api

import (
	"net/http"
)

func (s *server) routes() *http.ServeMux {
	router := s.router

	router.HandleFunc("/status", s.ApiStatus())
	router.HandleFunc("/rates/latest", s.GetLatest())
	router.HandleFunc("/rates/{date}", s.GetSpecificDate())
	router.HandleFunc("/rates/analyze", s.AnalyzeRates())
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
