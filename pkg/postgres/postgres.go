package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgreDB Хранилище данных
type PostgreDB struct {
	PG *gorm.DB
}

// New Конструктор
func New(constr string) (*PostgreDB, error) {
	// Инициализация GORM
	gormDB, err := gorm.Open(postgres.Open(constr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	s := PostgreDB{
		PG: gormDB,
	}
	return &s, nil
}
