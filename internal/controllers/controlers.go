package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zatrasz75/task_junior/internal/models"
	"github.com/zatrasz75/task_junior/pkg/logger"
	"github.com/zatrasz75/task_junior/pkg/storage"
	"net/http"
	"strconv"
	"sync"
)

type Server struct {
	PG storage.Database
}

// ReceiveSave Метод для обработки POST-запроса с JSON-данными
func (s *Server) ReceiveSave(w http.ResponseWriter, r *http.Request) {
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
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		age, err := getAgeFromAPI(p.Name)
		if err != nil {
			http.Error(w, "Не удалось получить данные о возрасте", http.StatusInternalServerError)
			logger.Error("Не удалось получить данные о возрасте", err)
		}
		p.Age = age
	}()
	go func() {
		defer wg.Done()
		gender, err := getGenderFromAPI(p.Name)
		if err != nil {
			http.Error(w, "Не удалось получить данные о поле", http.StatusInternalServerError)
			logger.Error("Не удалось получить данные о поле", err)
		}
		p.Gender = gender
	}()
	go func() {
		defer wg.Done()
		nationality, err := getNationalityFromAPI(p.Name)
		if err != nil {
			http.Error(w, "Не удалось получить данные о национальности", http.StatusInternalServerError)
			logger.Error("Не удалось получить данные о национальности", err)
		}
		p.Nationality = nationality
	}()

	wg.Wait()

	id, err := s.PG.SavePersonToDate(p)
	if err != nil {
		logger.Error("не получилось сохранить данные в базу данных", err)
		return
	}
	logger.Info("Данные успешно сохранены, получен номер id %d записи", id)

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

// GetData Метод для обработки GET-запроса с JSON-данными
func (s *Server) GetData(w http.ResponseWriter, r *http.Request) {
	genderFilter := r.URL.Query().Get("gender")  // фильтрация данных по полу.
	pageStr := r.URL.Query().Get("page")         // текущая страница данных.
	pageSizeStr := r.URL.Query().Get("pageSize") //  количество записей.
	ageStr := r.URL.Query().Get("age")           // возраст

	page := parseQueryParam(pageStr)
	pageSize := parseQueryParam(pageSizeStr)
	age, err := strconv.Atoi(ageStr)
	if err != nil || age <= 0 {
		age = 0
	}

	data, err := s.PG.Select(genderFilter, age, page, pageSize)
	if err != nil {
		logger.Error("Ошибка при выполнении запроса к базе данных", err)
		http.Error(w, "Ошибка при выполнении запроса к базе данных", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (s *Server) DeleteData(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	if idParam == "" {
		logger.Info("Отсутствует идентификатор")
		http.Error(w, "Отсутствует идентификатор", http.StatusBadRequest)
		return
	}
	id := parseQueryParam(idParam)

	err := s.PG.DeleteDataByID(id)
	if err != nil {
		logger.Error("Ошибка при удалении данных", err)
		http.Error(w, "Ошибка при удалении данных", http.StatusInternalServerError)
		return
	}
	logger.Info("Данные c id %d успешно удалены", id)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Данные успешно удалены"}
	json.NewEncoder(w).Encode(response)
}
