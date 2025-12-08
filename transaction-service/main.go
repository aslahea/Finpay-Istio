package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Transaction struct {
	ID        string    `json:"id"`
	PaymentID string    `json:"payment_id"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	tx := Transaction{
		ID:        "t-9001",
		PaymentID: "p-5001",
		Status:    "RECORDED",
		Timestamp: time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tx)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/transaction", transactionHandler)
	http.HandleFunc("/healthz", healthHandler)

	log.Println("transaction-service listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
