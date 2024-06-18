package usecases

import (
	"github.com/blanc42/mooca/pkg/models"
	"github.com/blanc42/mooca/pkg/repository"
)

type CourseUsecase struct {
	repo repository.CourseRepositoryInterface
}

func NewCourseUsecase(repo repository.CourseRepositoryInterface) CourseUsecaseInterface {
	return &CourseUsecase{repo: repo}
}

func (uc *CourseUsecase) CreateCourse(course *models.Course) error {
	return uc.repo.Create(course)
}

func (uc *CourseUsecase) UpdateCourse(course *models.Course) error {
	return uc.repo.Update(course)
}

func (uc *CourseUsecase) DeleteCourse(course *models.Course) error {
	return uc.repo.Delete(course)
}

func (uc *CourseUsecase) GetCourseByID(id uint) (*models.Course, error) {
	return uc.repo.FindByID(id)
}

func (uc *CourseUsecase) GetAllCourses() ([]models.Course, error) {
	return uc.repo.GetAll()
}
