package service

import (
	"edu_test_graph/graph/model"
	"edu_test_graph/internal/repository"
)

type QuestionService struct {
	repo *repository.QuestionRepository
}

func NewQuestionService(repo *repository.QuestionRepository) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) CreateQuestion(question *model.Question) error {
	// Business logic for creating a question
	return s.repo.Create(question)
}

func (s *QuestionService) GetQuestion(id int) (*model.Question, error) {
	// Business logic for retrieving a question
	return s.repo.Get(id)
}

// Define other methods as needed
