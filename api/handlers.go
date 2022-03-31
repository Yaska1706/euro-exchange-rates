package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

const (
	EUROfxREF = "EUR"
)

func (s *server) ApiStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			notFound(w, r)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := Response{
			Status: "success",
			Data:   "Api running smoothly",
		}
		byteres, _ := json.Marshal(response)
		w.Write(byteres)
	}
}

func (s *server) GetLatest() http.HandlerFunc {

	rates := []map[string]string{}
	latestrates := returnlatestrates()
	for _, latestrate := range latestrates {
		rates = append(rates, map[string]string{
			latestrate.Currency: latestrate.Rate,
		})
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			notFound(w, r)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{
			"base":  EUROfxREF,
			"rates": rates,
		}
		byteresp, _ := json.Marshal(response)

		w.Write(byteresp)
	}

}

func (s *server) GetSpecificDate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			notFound(w, r)
			return
		}
		params := mux.Vars(r)

		date := params["date"]

		rates := map[string]string{}

		latestrates := returnratesperdate(date)
		for _, latestrate := range latestrates {
			rates[latestrate.Currency] = latestrate.Rate
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := map[string]interface{}{
			"base":  EUROfxREF,
			"rates": rates,
		}
		byteresp, _ := json.Marshal(response)

		w.Write(byteresp)

	}
}

func (s *server) AnalyzeRates() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			notFound(w, r)
			return
		}

	}

}
