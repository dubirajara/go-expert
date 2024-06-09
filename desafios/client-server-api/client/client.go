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

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	switch {
	case err != nil:
		panic(err)
	case resp.StatusCode != http.StatusOK:
		panic(fmt.Errorf("%s status code: %d", string(body), resp.StatusCode))
	}

	var data server.QuoteCurrency
	if err := json.Unmarshal(body, &data); err != nil {
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
