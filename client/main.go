package main

import (
	"context"
	"io"
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

	// context que vai ser cancelado se não houver uma resposta em até 3 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)

	// req, err := http.Get("http://localhost:8080/cotacao")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
	// }
	// defer req.Body.Close()

	// res, err := io.ReadAll(req.Body)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Erro ao ler a resposta da requisição: %v\n", err)
	// }

	// var data QuotationAPI
	// err = json.Unmarshal(res, &data)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v\n", err)
	// }

	// f, err := os.Create("cotacao.txt")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Erro ao criar arquivo %v\n", err)
	// }
	// defer f.Close()

	// _, err = f.WriteString(fmt.Sprintf("Dólar {%s}", data.Usdbrl.Bid))
	// fmt.Println("Arquivo criado com sucesso!")
	// fmt.Println("Dólar: ", data.Usdbrl.Bid)
}
