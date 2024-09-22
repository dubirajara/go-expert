package service

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/valyala/fastjson"
)

type WeatherData struct {
	Celsius   float64 `json:"temp_C"`
	Farenheit float64 `json:"temp_F"`
	Kelvin    float64 `json:"temp_K"`
}

func GetWeather(city, apiKey string) (*WeatherData, error) {
	baseURL := "https://api.weatherapi.com/v1/current.json"
	params := url.Values{}
	params.Add("q", city)
	params.Add("key", apiKey)
	encodedParams := params.Encode()
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?%s", baseURL, encodedParams), nil)
	if err != nil {
		return nil, err
	}

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 401 || res.StatusCode == 403 {
		return nil, errors.New("API key is invalid or not provided")
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

	currentWeather := v.Get("current")
	tempC := currentWeather.GetFloat64("temp_c")
	tempF := currentWeather.GetFloat64("temp_f")
	tempK := tempC + 273
	data := WeatherData{
		Celsius:   tempC,
		Farenheit: tempF,
		Kelvin:    tempK,
	}
	return &data, nil
}
