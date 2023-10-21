package controllers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zatrasz75/task_junior/pkg/logger"
	"net/http"
	"strconv"
)

func getAgeFromAPI(name string) (int, error) {
	apiUrl := fmt.Sprintf("https://api.agify.io/?name=" + name)

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := httpClient.Get(apiUrl)
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
		logger.Info("Обогащен возраст %d ", *result.Age)
		return *result.Age, nil
	}

	return 0, nil
}

func getGenderFromAPI(name string) (string, error) {
	apiUrl := fmt.Sprintf("https://api.genderize.io/?name=" + name)

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := httpClient.Get(apiUrl)
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
	logger.Info("Обогащен пол %s ", result.Gender)

	return result.Gender, nil
}

func getNationalityFromAPI(name string) (string, error) {
	apiUrl := fmt.Sprintf("https://api.nationalize.io/?name=" + name)

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := httpClient.Get(apiUrl)
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
	logger.Info("Обогащена национальность %s ", nationality)

	return nationality, nil
}

// Вспомогательная функция для преобразования строки в число .
func parseQueryParam(param string) int {
	value, err := strconv.Atoi(param)
	if err != nil || value <= 0 {
		// Если произошла ошибка или значение некорректное, используем значение по умолчанию.
		return 1
	}
	return value
}
