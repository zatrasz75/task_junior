package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

func getAgeFromAPI(name string) (int, error) {
	apiUrl := fmt.Sprintf("https://api.agify.io/?name=" + name)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return 0, errors.New("Не удалось получить данные о возрасте")
	}

	var result struct {
		Age *int `json:"age"`
	}

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&result); err != nil {
		return 0, err
	}
	if result.Age != nil {
		return *result.Age, nil
	}

	return 0, nil
}

func getGenderFromAPI(name string) (string, error) {
	apiUrl := fmt.Sprintf("https://api.genderize.io/?name=" + name)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return "", errors.New("Не удалось получить данные о возрасте")
	}

	var result struct {
		Gender string `json:"gender"`
	}

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&result); err != nil {
		return "", err
	}

	return result.Gender, nil
}

func getNationalityFromAPI(name string) (string, error) {
	apiUrl := fmt.Sprintf("https://api.nationalize.io/?name=" + name)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return "", errors.New("Не удалось получить данные о национальности")
	}

	var result struct {
		Country []struct {
			CountryID   string  `json:"country_id"`
			Probability float64 `json:"probability"`
		} `json:"country"`
	}

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&result); err != nil {
		return "", err
	}

	// Найдем страну с максимальной вероятностью
	maxProbability := 0.0
	var nationality string

	for _, country := range result.Country {
		if country.Probability > maxProbability {
			maxProbability = country.Probability
			nationality = country.CountryID
		}
	}

	return nationality, nil
}
