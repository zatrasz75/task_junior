package storage

import "github.com/zatrasz75/task_junior/internal/models"

type Database interface {
	SavePersonToDate(p models.Person) (int, error)
	Select(genderFilter string, page, pageSize, size int) ([]models.Person, error)
	DeleteDataByID(id int) error
}
