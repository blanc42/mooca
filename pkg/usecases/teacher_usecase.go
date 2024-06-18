package usecases

import "github.com/blanc42/mooca/pkg/models"

type TeacherUsecaseInterface interface {
	CreateTeacher(teacher *models.Teacher) error
	UpdateTeacher(teacher *models.Teacher) error
	DeleteTeacher(teacherID uint) error
	GetTeacherByID(id uint) (*models.Teacher, error)
	GetAllTeachers() ([]models.Teacher, error)
}
