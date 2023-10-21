package storage

import "github.com/zatrasz75/task_junior/internal/models"

type Database interface {
	Select(genderFilter string, page, pageSize, size int) ([]models.Person, error)
	DeleteDataByID(id int) error
	UpdateDataByID(id int, newData models.Person) error
	PartialUpdateDataByID(id int, partialData map[string]interface{}) error
	SaveNewPeople(newData models.Person) (int, error)
}
