package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type PaymentResponse struct {
	ID        string    `json:"id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	payment := PaymentResponse{
		ID:        "p-5001",
		Amount:    199.99,
		Status:    "INITIATED",
		Timestamp: time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/payment", paymentHandler)
	http.HandleFunc("/healthz", healthHandler)

	log.Println("payment-service v1 listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
