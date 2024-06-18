package usecases

import (
	"github.com/blanc42/mooca/pkg/models"
	"github.com/blanc42/mooca/pkg/repository"
)

type TeacherUsecase struct {
	repo repository.TeacherRepositoryInterface
}

func NewTeacherUsecase(repo repository.TeacherRepositoryInterface) *TeacherUsecase {
	return &TeacherUsecase{repo: repo}
}

func (uc *TeacherUsecase) CreateTeacher(teacher *models.Teacher) error {
	return uc.repo.Create(teacher)
}

func (uc *TeacherUsecase) UpdateTeacher(teacher *models.Teacher) error {
	return uc.repo.Update(teacher)
}

func (uc *TeacherUsecase) DeleteTeacher(teacherId uint) error {
	return uc.repo.Delete(teacherId)
}

func (uc *TeacherUsecase) GetTeacherByID(id uint) (*models.Teacher, error) {
	return uc.repo.FindByID(id)
}

func (uc *TeacherUsecase) GetAllTeachers() ([]models.Teacher, error) {
	return uc.repo.GetAll()
}
