package repository

import "github.com/blanc42/mooca/pkg/models"

type InstituteRepository interface {
	CreateInstitute(institute *models.Institute) error
	GetInstituteByID(id uint) (*models.Institute, error)
	GetAllInstitutes() ([]models.Institute, error)
	UpdateInstitute(institute *models.Institute) error
	DeleteInstitute(id uint) error
}
