package usecases

import "github.com/blanc42/mooca/pkg/models"

type CourseUsecaseInterface interface {
	CreateCourse(course *models.Course) error
	UpdateCourse(course *models.Course) error
	DeleteCourse(course *models.Course) error
	GetCourseByID(id uint) (*models.Course, error)
	GetAllCourses() ([]models.Course, error)
}
