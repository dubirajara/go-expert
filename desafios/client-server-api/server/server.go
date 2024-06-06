package server

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/valyala/fastjson"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type QuoteCurrency struct {
	Code       string `gorm:"primaryKey" json:"code"`
	Codein     string `gorm:"primaryKey" json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func Start() {
	log.Println("Server started...")
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", QuoteCurrencyHandler)
	if error := http.ListenAndServe(":8080", mux); error != nil {
		log.Fatalf("Could not listen th server: %v\n", error)
	}

}

func QuoteCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started...")
	resp, err := getQuoteCurrency()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer log.Println("Request finished.")
	select {
	case <-time.After(200 * time.Millisecond):
		log.Println("Request processed successfully")
		go SaveDBQuoteCurrency(resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)

	case <-ctx.Done():
		msg := "Request canceled by client"
		log.Println(msg)
		http.Error(w, msg, http.StatusRequestTimeout)

	}
}

func getQuoteCurrency() (*QuoteCurrency, error) {
	req, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")

	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	resp, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var p fastjson.Parser
	v, err := p.Parse(string(resp))
	if err != nil {
		return nil, err
	}

	var data QuoteCurrency
	if err := json.Unmarshal([]byte(v.Get("USDBRL").String()), &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func SaveDBQuoteCurrency(data *QuoteCurrency) {
	log.Println("Saving data to database...")
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&QuoteCurrency{})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "code"}, {Name: "codein"}},
		DoUpdates: clause.AssignmentColumns([]string{"bid"}),
	}).Create(&data)

	select {

	case <-ctx.Done():
		log.Println("Timeout reached")

	default:
		log.Println("Data saved successfully")

	}

}
