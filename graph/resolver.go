package graph

import "edu_test_graph/internal/service"

type Resolver struct {
	GroupService      *service.GroupService
	AnswerService     *service.AnswerService
	CollectionService *service.CollectionService
	StudentService    *service.StudentService
	QuestionService   *service.QuestionService
}
