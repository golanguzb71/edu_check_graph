package repository

import (
	"context"
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
	_, err := r.db.Exec("INSERT INTO questions (question_field, collection_id) VALUES ($1, $2)",
		question.QuestionField, question.CollectionID)
	if err != nil {
		log.Printf("Error inserting question: %v", err)
		return err
	}
	return nil
}

// Get retrieves a question by ID.
func (r *QuestionRepository) Get(id int) (*model.Question, error) {
	row := r.db.QueryRow("SELECT id, question_field, collection_id, created_at, updated_at FROM questions WHERE id = $1", id)
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
	_, err := r.db.Exec("UPDATE questions SET question_field = $1, collection_id = $2, updated_at = $3 WHERE id = $4",
		question.QuestionField, question.CollectionID, question.UpdatedAt, question.ID)
	if err != nil {
		log.Printf("Error updating question: %v", err)
		return err
	}
	return nil
}

// Delete removes a question from the database.
func (r *QuestionRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM questions WHERE id = $1", id)
	if err != nil {
		log.Printf("Error deleting question: %v", err)
		return err
	}
	return nil
}

func (r *QuestionRepository) CreateTest(ctx context.Context, collectionId string, questions []*model.TestQuestion) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	for _, question := range questions {
		var questionId int
		err = tx.QueryRowContext(ctx, `INSERT INTO questions (question_field, collection_id)
			VALUES ($1, $2)
			RETURNING id`, question.QuestionField, collectionId).Scan(&questionId)
		if err != nil {
			return err
		}

		for _, answer := range question.Answers {
			_, err = tx.ExecContext(ctx, `INSERT INTO answers (is_true, question_id, answer_field)
				VALUES ($1, $2, $3)`, answer.IsTrue, questionId, answer.AnswerField)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
