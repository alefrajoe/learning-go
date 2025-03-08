package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TimeJson struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth string `json:"day_of_month"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Hour       string `json:"hour"`
	Minute     string `json:"minute"`
	Second     string `json:"second"`
}

func main() {
	mux := http.NewServeMux()
	// Update the function call to use the imported middleware
	mux.Handle("/", loggingIncomingIP(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Accept") == "application/json" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(TimeJson{
				DayOfWeek:  time.Now().Format("Monday"),
				DayOfMonth: time.Now().Format("2"),
				Month:      time.Now().Format("January"),
				Year:       time.Now().Format("06"),
				Hour:       time.Now().Format("15"),
				Minute:     time.Now().Format("04"),
				Second:     time.Now().Format("05"),
			})
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(time.Now().Format(time.RFC3339)))
		}
	})))
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
