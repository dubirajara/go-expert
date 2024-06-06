package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"server/server"
	"time"
)

func main() {
	ch := make(chan string)
	go client(ch)
	result := <-ch
	saveQuoteBid(result)

}

func client(ch chan string) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
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

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var data server.QuoteCurrency
	if err := json.Unmarshal(resp, &data); err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stdout, data.Bid)
	ch <- data.Bid

}

func saveQuoteBid(bid string) {
	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("Dólar: " + bid)

}
