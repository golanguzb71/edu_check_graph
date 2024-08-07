package service

import (
	"edu_test_graph/graph/model"
	"edu_test_graph/internal/repository"
	"strconv"
)

type CollectionService struct {
	repo *repository.CollectionRepository
}

func NewCollectionService(repo *repository.CollectionRepository) *CollectionService {
	return &CollectionService{repo: repo}
}

func (s *CollectionService) CreateCollection(collection *model.Collection) error {
	return s.repo.Create(collection)
}

func (s *CollectionService) GetCollection(id int) (*model.FullCollection, error) {
	return s.repo.Get(id)
}

func (s *CollectionService) UpdateCollection(m *model.Collection) error {
	return s.repo.Update(m)
}

func (s *CollectionService) DeleteCollection(id string) error {
	realId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(realId)
}

func (s *CollectionService) GetCollections() ([]*model.Collection, error) {
	return s.repo.GetCollections()
}
