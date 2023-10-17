package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/raphapaulino/pos-graduacao-goexpert-desafio-1-client-server-api/types"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
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

type QuotationDB struct {
	ID  string
	Bid string
}

func NewQuotationDB(bid string) *QuotationDB {
	return &QuotationDB{
		ID:  uuid.NewString(),
		Bid: bid,
	}
}

// VERIFICAR A FUNÇÃO MAIN DESTE EXEMPLO: https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go
func main() {
	http.HandleFunc("/cotacao", searchQuotationHandler)
	http.ListenAndServe(":8080", nil)
}

func searchQuotationHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Iniciou o método: searchQuotationHandler")

	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	r.Header.Set("Accept", "application/json")

	// // throws HTTP 404 if the url path were not http://localhost:8080/cotacao
	// if r.URL.Path != "/cotacao" {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }

	quotation, err := SearchQuotation(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	payload := types.QuotationResponse{BID: quotation.Usdbrl.Bid}

	json.NewEncoder(w).Encode(payload)

	// abre a conexão com o banco de dados
	db, err := sql.Open("sqlite3", "./posgosqlite-2.db")
	checkErr(err)
	defer db.Close()

	// gera a nova cotação
	quotation2 := NewQuotationDB(quotation.Usdbrl.Bid)
	err = insertQuotationDB(db, quotation2)
	checkErr(err)

	// ctx := r.Context()
	// log.Println("Request started")
	// successMsg := "Request processed with success"
	// cancelMsg := "Request canceled by the client"
	// defer log.Println("Request finished")

	// select {
	// 	case <-time.After(5 * time.Second):
	// 		// imprime no command line stdout
	// 		log.Println(successMsg)
	// 		// imprime no browser
	// 		w.Write([]byte(successMsg))

	// 	case <-ctx.Done():
	// 		// imprime no command line stdout
	// 		log.Println(cancelMsg)
	// }
}

func SearchQuotation(ctx context.Context) (*QuotationAPI, error) {

	log.Println("Iniciou o método: SearchQuotation")
	// var cotacao map[string]types.QuotationDataDTO

	// ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	// defer cancel()

	// req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, types.QUOTATION_API_URL, nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// // exibe o response da chamada get da api na command line
	// io.Copy(os.Stdout, res.Body)

	body, error := ioutil.ReadAll(res.Body)
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

func insertQuotationDB(db *sql.DB, quotation *QuotationDB) error {
	stmt, err := db.Prepare("INSERT INTO quotations(id, bid) VALUES(?, ?)")
	// stmt, err := db.Prepare("INSERT INTO quotations(id, bid) VALUES($1, $2)") // para o sqlite
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(quotation.ID, quotation.Bid)
	if err != nil {
		return err
	}

	fmt.Printf("Criado o registro do Bid no valor de: %v \n", quotation.Bid)

	return nil
}

func checkErr(err error) {
	if err != nil {
		// panic(err)
		log.Fatal(err)
	}
}
