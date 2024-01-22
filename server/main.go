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

// VERIFICAR A FUNÇÃO MAIN DESTE EXEMPLO: https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go
func main() {
	http.HandleFunc("/cotacao", searchQuotationHandler)
	http.ListenAndServe(":8080", nil)
}

func searchQuotationHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Iniciou o método: searchQuotationHandler")
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	defer log.Println("Request finalizada")

	r.Header.Set("Accept", "application/json")

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
}

func SearchQuotation(ctx context.Context) (*QuotationAPI, error) {

	log.Println("Iniciou o método: SearchQuotation...")

	// req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, types.QUOTATION_API_URL, nil)
	if err != nil {
		checkErr(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		checkErr(err)
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
	log.Println("Terminou o método: SearchQuotation")

	return &q, nil
}

func NewQuotationDB(bid string) *QuotationDB {
	return &QuotationDB{
		ID:  uuid.NewString(),
		Bid: bid,
	}
}

func insertQuotationDB(db *sql.DB, quotation *QuotationDB) error {
	log.Println("Iniciou o método: insertQuotationDB...")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond) // time.Microsecond
	defer cancel()

	stmt, err := db.Prepare("INSERT INTO quotations(id, bid) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, quotation.ID, quotation.Bid)
	if err != nil {
		checkErr(err)
	}

	fmt.Printf("Criado o registro do Bid no valor de: %v \n", quotation.Bid)
	log.Println("Terminou o método: insertQuotationDB")

	return nil
}

func checkErr(err error) {
	if err != nil {
		// panic(err)
		log.Fatal(err)
	}
}
