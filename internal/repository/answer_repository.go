package repository

import (
	"database/sql"
	"edu_test_graph/graph/model"
	"log"
)

// AnswerRepository handles operations related to the Answer model.
type AnswerRepository struct {
	db *sql.DB
}

func NewAnswerRepository(db *sql.DB) *AnswerRepository {
	return &AnswerRepository{db: db}
}

// Create inserts a new answer into the database.
func (r *AnswerRepository) Create(answer *model.Answer) error {
	_, err := r.db.Exec("INSERT INTO answers (is_true, question_id, answer_field) VALUES ($1, $2, $3)",
		answer.IsTrue, answer.QuestionID, answer.AnswerField)
	if err != nil {
		log.Printf("Error inserting answer: %v", err)
		return err
	}
	return nil
}

// Get retrieves an answer by ID.
func (r *AnswerRepository) Get(id int) (*model.Answer, error) {
	row := r.db.QueryRow("SELECT id, is_true, question_id, answer_field, created_at, updated_at FROM answers WHERE id = $1", id)
	answer := &model.Answer{}
	err := row.Scan(&answer.ID, &answer.IsTrue, &answer.QuestionID, &answer.AnswerField, &answer.CreatedAt, &answer.UpdatedAt)
	if err != nil {
		log.Printf("Error retrieving answer: %v", err)
		return nil, err
	}
	return answer, nil
}

// Update updates an existing answer in the database.
func (r *AnswerRepository) Update(answer *model.Answer) error {
	_, err := r.db.Exec("UPDATE answers SET is_true = $1, question_id = $2, answer_field = $3, updated_at = $4 WHERE id = $5",
		answer.IsTrue, answer.QuestionID, answer.AnswerField, answer.UpdatedAt, answer.ID)
	if err != nil {
		log.Printf("Error updating answer: %v", err)
		return err
	}
	return nil
}

func (r *AnswerRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM answers WHERE id = $1", id)
	if err != nil {
		log.Printf("Error deleting answer: %v", err)
		return err
	}
	return nil
}
