package service

import (
	"edu_test_graph/graph/model"
	"edu_test_graph/internal/repository"
)

type GroupService struct {
	repo *repository.GroupRepository
}

func NewGroupService(repo *repository.GroupRepository) *GroupService {
	return &GroupService{repo: repo}
}

func (s *GroupService) CreateGroup(group *model.Group) error {
	return s.repo.Create(group)
}

func (s *GroupService) GetGroup(id *string, orderLevel *bool) ([]*model.Group, error) {
	return s.repo.Get(id, orderLevel)
}

func (s *GroupService) UpdateGroup(group *model.Group) error {
	return s.repo.Update(group)
}

func (s *GroupService) DeleteGroup(id int) error {
	return s.repo.Delete(id)
}
