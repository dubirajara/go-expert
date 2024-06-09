package server

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"runtime/debug"
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

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Recovered panic", r)
				debug.PrintStack()
				http.Error(w, "Something went wrong, internal error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func Start() {
	log.Println("Server started...")
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", QuoteCurrencyHandler)
	if error := http.ListenAndServe(":8080", recoveryMiddleware(mux)); error != nil {
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

	case <-ctx.Done():
		msg := "Request canceled by client"
		log.Println(msg)
		http.Error(w, msg, http.StatusRequestTimeout)

	default:
		log.Println("Request processed successfully")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		go SaveDBQuoteCurrency(resp)

	}
}

func getQuoteCurrency() (*QuoteCurrency, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
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
