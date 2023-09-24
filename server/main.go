package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type QuotationAPI struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	http.HandleFunc("/cotacao", searchQuotationHandler)
	http.ListenAndServe(":8080", nil)
}

func searchQuotationHandler(w http.ResponseWriter, r *http.Request) {

	r.Header.Set("Accept", "application/json")

	// throws HTTP 404 if the url path were not http://localhost:8080/cotacao
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	q, err := SearchQuotation()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(q)

	// ctx := r.Context()
	// log.Println("Request started")
	// successMsg := "Request processed with success"
	// cancelMsg := "Request canceled by the client"
	// defer log.Println("Request finished")

	// select {
	// 	case <-time.After(5 * time.Second):
	// 		log.Println(successMsg)
	// 		w.Write([]byte(successMsg))

	// 	case <-ctx.Done():
	// 		log.Println(cancelMsg)
	// }
}

func SearchQuotation() (*QuotationAPI, error) {
	resp, error := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()

	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var q QuotationAPI
	error = json.Unmarshal(body, &q)
	if error != nil {
		return nil, error
	}

	return &q, nil
}
