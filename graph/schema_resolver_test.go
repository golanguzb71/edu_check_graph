package graph

import (
	"context"
	"edu_test_graph/graph/model"
	"testing"
)

func TestCreateTest(t *testing.T) {
	ctx := context.Background()
	collectionID := "1"
	questions := []*model.TestQuestion{
		{"Test field", []*model.AnswerInput{{true, "checking"}}},
		{"Test field", []*model.AnswerInput{{true, "checking"}}},
		{"Test field", []*model.AnswerInput{{true, "checking"}}},
	}

	r := &mutationResolver{}

	response, err := r.CreateTest(ctx, collectionID, questions)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if response == nil {
		t.Fatalf("expected response, got nil")
	}

	if response.Message != "Success" {
		t.Errorf("expected message 'Success', got %s", response.Message)
	}
}
