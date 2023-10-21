package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// PostgreDB Хранилище данных
type PostgreDB struct {
	PG *gorm.DB
	DB *pgxpool.Pool
}

// New Конструктор
func New(constr string) (*PostgreDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var db, err = pgxpool.Connect(ctx, constr)
	if err != nil {
		return nil, err
	}
	// Инициализация GORM
	gormDB, err := gorm.Open(postgres.Open(constr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	s := PostgreDB{
		PG: gormDB,
		DB: db,
	}
	return &s, nil
}
