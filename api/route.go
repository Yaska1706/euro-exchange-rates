package api

import (
	"database/sql"
	"net/http"
)

type Server struct {
	db *sql.DB
}

func (rh *rateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet:
		rh.GetLatest(w, r)
		return
	case r.Method == http.MethodGet:
		rh.GetSpecificDate(w, r)
		return
	case r.Method == http.MethodGet:
		rh.AnalyzeRates(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func Serve() {
	rateHandler := rateHandler{}
	mux := http.NewServeMux()
	mux.HandleFunc("/rates/latest", rateHandler.GetLatest)
	mux.HandleFunc("/rates/analyze", rateHandler.AnalyzeRates)
	http.ListenAndServe(":8080", mux)

}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}
