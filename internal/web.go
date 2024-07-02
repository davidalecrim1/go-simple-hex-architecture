package internal

import (
	"encoding/json"
	"log"
	"net/http"
)

type TaxWeb struct {
	svc *TaxService
}

// here passing by value not reference because don't want the
// same request to mess with the instance of another request
// (would happen using pointers).
func NewTaxWeb(svc TaxService) *TaxWeb {
	return &TaxWeb{svc: &svc}
}

type TaxRequest struct {
	Price float64 `json:"price"`
	Rate  float64 `json:"rate"`
}

type TaxResponse struct {
	Result float64 `json:"result"`
}

func (t *TaxWeb) TaxHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // Ensure the request body is closed after reading

	var reqBody TaxRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqBody)

	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := t.svc.Calculate(reqBody.Price, reqBody.Rate)

	if err != nil {
		log.Println(err)
		http.Error(w, "Unexpected error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(TaxResponse{Result: result})
}
