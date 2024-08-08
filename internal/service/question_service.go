package service

import (
	"context"
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
	return s.repo.Create(question)
}

func (s *QuestionService) GetQuestion(id int) (*model.Question, error) {
	// Business logic for retrieving a question
	return s.repo.Get(id)
}

func (s *QuestionService) CreateTest(ctx context.Context, id string, questions []*model.TestQuestion) error {
	return s.repo.CreateTest(ctx, id, questions)
}

func (s *QuestionService) DeleteQuestion(id int) error {
	return s.repo.Delete(id)
}
