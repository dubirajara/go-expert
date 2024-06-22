package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "CEP cannot be empty, try run command with CEP: go run main.go 78068-685\n")
	}
	for _, cep := range args {
		ch1 := make(chan string)
		ch2 := make(chan string)

		brasilApiUrl := "https://brasilapi.com.br/api/cep/v1/" + cep
		viaCepUrl := "https://viacep.com.br/ws/" + cep + "/json/"

		go GetCEPData(brasilApiUrl, ch1)
		go GetCEPData(viaCepUrl, ch2)

		select {
		case msg := <-ch1: // Brasilapi
			fmt.Printf("%v result: %v\n", brasilApiUrl, msg)
		case msg := <-ch2: // ViaCep
			fmt.Printf("%v result: %v\n", viaCepUrl, msg)

		case <-time.After(time.Second):
			fmt.Println("timeout")
		}
	}

}

func GetCEPData(url string, ch chan<- string) {
	req, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error to request: %v\n", err)
	}
	defer req.Body.Close()

	resp, err := io.ReadAll(req.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error to read response: %v\n", err)
	}

	ch <- string(resp)
}
