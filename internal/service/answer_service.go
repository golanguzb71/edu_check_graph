package service

import (
	"edu_test_graph/graph/model"
	"edu_test_graph/internal/repository"
)

type AnswerService struct {
	repo *repository.AnswerRepository
}

func NewAnswerService(repo *repository.AnswerRepository) *AnswerService {
	return &AnswerService{repo: repo}
}

func (s *AnswerService) CreateAnswer(answer *model.Answer) error {
	// Business logic for creating an answer
	return s.repo.Create(answer)
}

func (s *AnswerService) GetAnswer(id int) (*model.Answer, error) {
	// Business logic for retrieving an answer
	return s.repo.Get(id)
}

func (s *AnswerService) DeleteAnswer(id int) error {
	return s.repo.Delete(id)
}

func (s *AnswerService) InsertTestAnswer(answers model.AnswerInsert) (*model.CommonResponse, error) {
	return s.repo.CreateInsertAnswer(answers)
}
