package usecases

import (
	"github.com/blanc42/mooca/pkg/models"
	"github.com/blanc42/mooca/pkg/repository"
)

type instituteUseCase struct {
	instituteRepository repository.InstituteRepository
}

func NewInstituteUseCase(instituteRepository repository.InstituteRepository) InstituteUseCase {
	return &instituteUseCase{instituteRepository: instituteRepository}
}

func (u *instituteUseCase) CreateInstitute(institute *models.Institute) error {
	return u.instituteRepository.CreateInstitute(institute)
}

func (u *instituteUseCase) GetInstituteByID(id uint) (*models.Institute, error) {
	return u.instituteRepository.GetInstituteByID(id)
}

func (u *instituteUseCase) GetAllInstitutes() ([]models.Institute, error) {
	return u.instituteRepository.GetAllInstitutes()
}

func (u *instituteUseCase) UpdateInstitute(institute *models.Institute) error {
	return u.instituteRepository.UpdateInstitute(institute)
}

func (u *instituteUseCase) DeleteInstitute(id uint) error {
	return u.instituteRepository.DeleteInstitute(id)
}
