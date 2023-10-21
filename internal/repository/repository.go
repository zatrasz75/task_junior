package repository

import (
	"context"
	"fmt"
	"github.com/zatrasz75/task_junior/internal/models"
	"github.com/zatrasz75/task_junior/pkg/postgres"
)

type Store struct {
	*postgres.PostgreDB
}

func New(pg *postgres.PostgreDB) *Store {
	return &Store{pg}
}

// SavePersonToDate Сохраняет данные в базу данных и возвращает ID записи
func (s *Store) SavePersonToDate(p models.Person) (int, error) {
	result := s.PG.Create(&p)
	if result.Error != nil {
		return 0, result.Error
	}

	return p.ID, nil
}

// Select выполняет SQL-запрос для выборки данных из таблицы с фильтрами и пагинацией.
func (s *Store) Select(genderFilter string, age, page, pageSize int) ([]models.Person, error) {
	// Создаем SQL-запрос с учетом фильтра и пагинации.
	query := "SELECT * FROM people WHERE true"
	if genderFilter != "" {
		query += fmt.Sprintf(" AND gender = '%s'", genderFilter)
	}
	if age > 0 {
		query += fmt.Sprintf(" AND age = %d", age)
	}
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, (page-1)*pageSize)

	// Выполняем запрос к базе данных.
	rows, err := s.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Person
	for rows.Next() {
		var data models.Person
		if err = rows.Scan(&data.ID, &data.Name, &data.Surname, &data.Patronymic, &data.Age, &data.Gender, &data.Nationality); err != nil {
			return nil, err
		}
		result = append(result, data)
	}

	return result, nil
}

// DeleteDataByID Создает SQL-запрос для удаления данных по идентификатору.
func (s *Store) DeleteDataByID(id int) error {
	delet := "DELETE FROM people WHERE id = $1"

	_, err := s.DB.Exec(context.Background(), delet, id)
	if err != nil {
		return err
	}

	return nil
}
