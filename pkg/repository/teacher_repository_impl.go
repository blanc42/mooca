package repository

import (
	"github.com/blanc42/mooca/pkg/models"
	"gorm.io/gorm"
)

type TeacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) *TeacherRepository {
	return &TeacherRepository{db: db}
}

func (repo *TeacherRepository) Create(teacher *models.Teacher) error {
	return repo.db.Create(teacher).Error
}

func (repo *TeacherRepository) Update(teacher *models.Teacher) error {
	return repo.db.Save(teacher).Error
}

func (repo *TeacherRepository) Delete(teacherId uint) error {
	return repo.db.Delete(&models.Teacher{Id: teacherId}).Error
}

func (repo *TeacherRepository) FindByID(id uint) (*models.Teacher, error) {
	var teacher models.Teacher
	result := repo.db.First(&teacher, id)
	return &teacher, result.Error
}

func (repo *TeacherRepository) GetAll() ([]models.Teacher, error) {
	var teachers []models.Teacher
	result := repo.db.Find(&teachers)
	return teachers, result.Error
}
