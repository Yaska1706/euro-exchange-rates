package api

import (
	"net/http"
	"sync"
)

type ExchangeRate struct {
	Date     string `json:"date"`
	Currency string `json:"currency"`
	Rate     string `json:"rate"`
}

type rateHandler struct {
	*sync.Mutex
}

func (rh *rateHandler) GetLatest(w http.ResponseWriter, r *http.Request) {

}

func (rh *rateHandler) GetSpecificDate(w http.ResponseWriter, r *http.Request) {

}

func (rh *rateHandler) AnalyzeRates(w http.ResponseWriter, r *http.Request) {

}
