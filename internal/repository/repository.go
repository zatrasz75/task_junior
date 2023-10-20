package repository

import (
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
