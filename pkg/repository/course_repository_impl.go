package repository

import (
	"github.com/blanc42/mooca/pkg/models"
	"gorm.io/gorm"
)

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepositoryInterface {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) Create(course *models.Course) error {
	return r.db.Create(course).Error
}

func (r *CourseRepository) Update(course *models.Course) error {
	return r.db.Save(course).Error
}

func (r *CourseRepository) Delete(course *models.Course) error {
	return r.db.Delete(course).Error
}

func (r *CourseRepository) FindByID(id uint) (*models.Course, error) {
	var course models.Course
	if err := r.db.First(&course, id).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *CourseRepository) GetAll() ([]models.Course, error) {
	var courses []models.Course
	if err := r.db.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}
