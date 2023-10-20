package storage

import "github.com/zatrasz75/task_junior/internal/models"

type Database interface {
	SavePersonToDate(p models.Person) (int, error)
}
