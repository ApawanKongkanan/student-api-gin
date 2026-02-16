package repositories

import (
	"errors"
	"go-api-gin/models"
	"gorm.io/gorm"
)

type StudentRepository interface {
	GetAll() ([]models.Student, error)
	GetByID(id string) (*models.Student, error)
	Create(student models.Student) (*models.Student, error)
	Update(student models.Student) (*models.Student, error)
	Delete(id string) error
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db}
}

func (r *studentRepository) GetAll() ([]models.Student, error) {
	var students []models.Student
	if err := r.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *studentRepository) GetByID(id string) (*models.Student, error) {
	var student models.Student
	if err := r.db.First(&student, "id = ?", id).Error; err != nil {
		return nil, errors.New("student not found")
	}
	return &student, nil
}

func (r *studentRepository) Create(student models.Student) (*models.Student, error) {
	if err := r.db.Create(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepository) Update(student models.Student) (*models.Student, error) {
	result := r.db.Model(&models.Student{}).
		Where("id = ?", student.ID).
		Updates(student)

	if result.RowsAffected == 0 {
		return nil, errors.New("student not found")
	}

	return &student, nil
}

func (r *studentRepository) Delete(id string) error {
	result := r.db.Delete(&models.Student{}, "id = ?", id)

	if result.RowsAffected == 0 {
		return errors.New("student not found")
	}

	return nil
}