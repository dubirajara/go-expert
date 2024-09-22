package service

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type ViaCEP struct {
	Localidade string `json:"localidade"`
	Erro       string `json:"erro,omitempty"`
}

func GetZipCode(zipCode string) (*ViaCEP, error) {
	req, err := http.NewRequest(http.MethodGet, "https://viacep.com.br/ws/"+zipCode+"/json/", nil)
	if err != nil {
		return nil, err
	}

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 400 {
		return nil, errors.New("invalid zipcode")
	}

	defer res.Body.Close()
	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data ViaCEP
	if err := json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	if data.Erro == "true" {
		return nil, errors.New("can not find zipcode")
	}
	return &data, nil
}
