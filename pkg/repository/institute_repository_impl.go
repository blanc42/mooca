package repository

import (
	"github.com/blanc42/mooca/pkg/models"
	"gorm.io/gorm"
)

type instituteRepository struct {
	db *gorm.DB
}

func NewInstituteRepository(db *gorm.DB) InstituteRepository {
	return &instituteRepository{db: db}
}

func (r *instituteRepository) CreateInstitute(institute *models.Institute) error {
	return r.db.Create(institute).Error
}

func (r *instituteRepository) GetInstituteByID(id uint) (*models.Institute, error) {
	var institute models.Institute
	if err := r.db.Preload("Courses").First(&institute, id).Error; err != nil {
		return nil, err
	}
	return &institute, nil
}

func (r *instituteRepository) GetAllInstitutes() ([]models.Institute, error) {
	var institutes []models.Institute
	if err := r.db.Preload("Courses").Find(&institutes).Error; err != nil {
		return nil, err
	}
	return institutes, nil
}

func (r *instituteRepository) UpdateInstitute(institute *models.Institute) error {
	return r.db.Save(institute).Error
}

func (r *instituteRepository) DeleteInstitute(id uint) error {
	return r.db.Delete(&models.Institute{}, id).Error
}
