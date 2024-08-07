package repository

import (
	"database/sql"
	"edu_test_graph/graph/model"
	"log"
)

// QuestionRepository handles operations related to the Question model.
type QuestionRepository struct {
	db *sql.DB
}

// NewQuestionRepository creates a new QuestionRepository.
func NewQuestionRepository(db *sql.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

// Create inserts a new question into the database.
func (r *QuestionRepository) Create(question *model.Question) error {
	_, err := r.db.Exec("INSERT INTO questions (question_field, collection_id) VALUES (?, ?)",
		question.QuestionField, question.CollectionID)
	if err != nil {
		log.Printf("Error inserting question: %v", err)
		return err
	}
	return nil
}

// Get retrieves a question by ID.
func (r *QuestionRepository) Get(id int) (*model.Question, error) {
	row := r.db.QueryRow("SELECT id, question_field, collection_id, created_at, updated_at FROM questions WHERE id = ?", id)
	question := &model.Question{}
	err := row.Scan(&question.ID, &question.QuestionField, &question.CollectionID, &question.CreatedAt, &question.UpdatedAt)
	if err != nil {
		log.Printf("Error retrieving question: %v", err)
		return nil, err
	}
	return question, nil
}

// Update updates an existing question in the database.
func (r *QuestionRepository) Update(question *model.Question) error {
	_, err := r.db.Exec("UPDATE questions SET question_field = ?, collection_id = ?, updated_at = ? WHERE id = ?",
		question.QuestionField, question.CollectionID, question.UpdatedAt, question.ID)
	if err != nil {
		log.Printf("Error updating question: %v", err)
		return err
	}
	return nil
}

// Delete removes a question from the database.
func (r *QuestionRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM questions WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting question: %v", err)
		return err
	}
	return nil
}
