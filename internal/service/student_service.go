package service

import (
	"edu_test_graph/graph/model"
	"edu_test_graph/internal/repository"
)

type StudentService struct {
	repo *repository.StudentRepository
}

func NewStudentService(repo *repository.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) CreateStudent(student *model.Student) error {
	// Business logic for creating a student
	return s.repo.Create(student)
}

func (s *StudentService) GetStudent(id int) (*model.Student, error) {
	// Business logic for retrieving a student
	return s.repo.Get(id)
}

// Define other methods as needed
