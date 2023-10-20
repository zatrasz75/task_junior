package controllers

import (
	"encoding/json"
	"github.com/zatrasz75/task_junior/internal/models"
	"github.com/zatrasz75/task_junior/pkg/logger"
	"github.com/zatrasz75/task_junior/pkg/storage"
	"net/http"
)

type Server struct {
	PG storage.Database
}

// GetPerson Метод для обработки POST-запроса с JSON-данными
func (s *Server) GetPerson(w http.ResponseWriter, r *http.Request) {
	var p models.Person

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		http.Error(w, "Failed to parse JSON request", http.StatusBadRequest)
		return
	}
	// Проверка наличия обязательных полей
	if p.Name == "" || p.Surname == "" {
		http.Error(w, "Отсутствует имя или фамилия в данных JSON", http.StatusBadRequest)
		logger.Debug("Отсутствует имя или фамилия в данных JSON")
		return
	}

	// Обогащаем данные из внешнего API
	age, err := getAgeFromAPI(p.Name)
	if err != nil {
		http.Error(w, "Не удалось получить данные о возрасте", http.StatusInternalServerError)
		logger.Error("Не удалось получить данные о возрасте", err)
		//return
	}
	gender, err := getGenderFromAPI(p.Name)
	if err != nil {
		http.Error(w, "Не удалось получить данные о поле", http.StatusInternalServerError)
		logger.Error("Не удалось получить данные о поле", err)
		//	return
	}
	nationality, err := getNationalityFromAPI(p.Name)
	if err != nil {
		http.Error(w, "Не удалось получить данные о национальности", http.StatusInternalServerError)
		logger.Error("Не удалось получить данные о национальности", err)
		//return
	}

	p.Age = age
	p.Gender = gender
	p.Nationality = nationality

	id, err := s.PG.SavePersonToDate(p)
	if err != nil {
		return
	}

	response := map[string]int{"id": id}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Ошибка при формировании JSON-ответа", http.StatusInternalServerError)
		logger.Error("Ошибка при формировании JSON-ответа", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
