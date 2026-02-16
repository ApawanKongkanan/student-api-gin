package services

import (
	"go-api-gin/models"
	"go-api-gin/repositories"
)

type StudentService interface {
	GetAll() ([]models.Student, error)
	GetByID(id string) (*models.Student, error)
	Create(student models.Student) (*models.Student, error)
	Update(student models.Student) (*models.Student, error)
	Delete(id string) error
}

type studentService struct {
	repo repositories.StudentRepository
}

func NewStudentService(repo repositories.StudentRepository) StudentService {
	return &studentService{repo}
}

func (s *studentService) GetAll() ([]models.Student, error) {
	return s.repo.GetAll()
}

func (s *studentService) GetByID(id string) (*models.Student, error) {
	return s.repo.GetByID(id)
}

func (s *studentService) Create(student models.Student) (*models.Student, error) {
	return s.repo.Create(student)
}

func (s *studentService) Update(student models.Student) (*models.Student, error) {
	return s.repo.Update(student)
}

func (s *studentService) Delete(id string) error {
	return s.repo.Delete(id)
}