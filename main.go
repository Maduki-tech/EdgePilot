package main

import (
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func main() {

	logfile := "logs/access.log"
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
		var message []byte
		_, err := r.Body.Read(message)

		if err != nil {
			logger.Println(err)
			return
		}

		logger.Println(message)

	}
}
