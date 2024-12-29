package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

var logger *log.Logger

type request struct {
	Name string
}

func main() {
	dateTime := time.Now().Format("2006-01-02T15:04:05")
	logfile := "logs/" + dateTime + ".log"

	logIo, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logIo.Close()

	logger = log.New(logIo, "", log.Ldate|log.Ltime)

	server := http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}

	http.HandleFunc("/", httpHandler)

	server.ListenAndServe()
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	// GET REQUEST
	if r.Method == http.MethodGet {

	}

	// POST REQUEST
	if r.Method == http.MethodPost {
		var req request
		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Println(req)

	}
}
