package repository

import "github.com/blanc42/mooca/pkg/models"

type TeacherRepositoryInterface interface {
	Create(teacher *models.Teacher) error
	Update(teacher *models.Teacher) error
	Delete(teacherId uint) error
	FindByID(id uint) (*models.Teacher, error)
	GetAll() ([]models.Teacher, error)
}
