package webserver

import (
	"cloudrun/internal/infra/service"
	"encoding/json"
	"log"
	"net/http"
)

type ErrorWeather struct {
	Message string `json:"message"`
}

func WeatherZipCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/weather/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	zipCode := r.URL.Query().Get("zipcode")

	respZipCode, err := service.GetZipCode(zipCode)
	if err != nil {
		error := err.Error()
		if error == "can not find zipcode" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorWeather{Message: error})
			return

		} else if error == "invalid zipcode" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(ErrorWeather{Message: error})
			return

		}
		log.Println(error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	apiKey := r.Context().Value("apiKey").(string)
	respWeather, err := service.GetWeather(respZipCode.Localidade, apiKey)
	if err != nil {
		error := err.Error()
		if error == "API key is invalid or not provided" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(ErrorWeather{Message: error})
			return

		}
		log.Println(error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(respWeather); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
