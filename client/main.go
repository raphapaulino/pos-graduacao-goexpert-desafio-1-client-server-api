package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/raphapaulino/pos-graduacao-goexpert-desafio-1-client-server-api/types"
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
	defer cancel()
	start := time.Now()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler a resposta da requisição: %v\n", err)
		panic(err)
	}

	// lê o response da requisição
	body, err := io.ReadAll(res.Body)
	checkErr(err)
	defer res.Body.Close()
	// io.Copy(os.Stdout, res.Body)

	// decodifica o json
	var data types.QuotationResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v\n", err)
	}

	// registra o resultado no arquivo cotacao.txt
	f, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar arquivo %v\n", err)
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("Dólar {%s} \n", data.BID))
	checkErr(err)
	fmt.Println("Arquivo criado com sucesso!")
	fmt.Fprintf(os.Stdout, "\nCotação do Dólar atual: {%v} \n", data.BID)

	elapsed := time.Since(start)
	log.Printf("The request took %s \n", elapsed)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
