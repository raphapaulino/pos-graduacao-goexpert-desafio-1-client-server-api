package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started")
	successMsg := "Request processed with success" 
	cancelMsg := "Request canceled by the client" 
	defer log.Println("Request finished")

	select {
		case <-time.After(5 * time.Second):
			log.Println(successMsg)
			w.Write([]byte(successMsg))

		case <-ctx.Done():
			log.Println(cancelMsg)
	}
}