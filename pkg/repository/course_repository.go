package repository

import "github.com/blanc42/mooca/pkg/models"

type CourseRepositoryInterface interface {
	Create(course *models.Course) error
	Update(course *models.Course) error
	Delete(course *models.Course) error
	FindByID(id uint) (*models.Course, error)
	GetAll() ([]models.Course, error)
}
