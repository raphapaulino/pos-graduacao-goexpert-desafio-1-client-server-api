package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
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

	// context que vai ser cancelado se não houver uma resposta em até 300 milissegundos
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	// ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond)
	defer cancel()
	start := time.Now()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		checkErr(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		checkErr(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)

	elapsed := time.Since(start)
	log.Printf("The request took %s \n", elapsed) // 123.925328ms
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
