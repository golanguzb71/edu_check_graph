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
	// Business logic for creating a group
	return s.repo.Create(group)
}

func (s *GroupService) GetGroup(id int) (*model.Group, error) {
	// Business logic for retrieving a group
	return s.repo.Get(id)
}

func (s *GroupService) UpdateGroup(group *model.Group) error {
	// Business logic for updating a group
	return s.repo.Update(group)
}

func (s *GroupService) DeleteGroup(id int) error {
	// Business logic for deleting a group
	return s.repo.Delete(id)
}
