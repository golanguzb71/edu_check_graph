package repository

import (
	"database/sql"
	"edu_test_graph/graph/model"
	"fmt"
	"log"
	"time"
)

// CollectionRepository handles operations related to the Collection model.
type CollectionRepository struct {
	db *sql.DB
}

// NewCollectionRepository creates a new CollectionRepository.
func NewCollectionRepository(db *sql.DB) *CollectionRepository {
	return &CollectionRepository{db: db}
}

// Create inserts a new collection into the database.
func (r *CollectionRepository) Create(collection *model.Collection) error {
	_, err := r.db.Exec("INSERT INTO collections (name) VALUES (?)", collection.Name)
	if err != nil {
		log.Printf("Error inserting collection: %v", err)
		return err
	}
	return nil
}

func (r *CollectionRepository) Get(id int) (*model.FullCollection, error) {
	rows, err := r.db.Query(`SELECT c.id,
		       c.name,
		       c.created_at,
		       c.updated_at,
		       q.id,
		       q.question_field,
		       q.created_at,
		       a.id,
		       a.created_at,
		       a.is_true,
		       a.answer_field
		FROM collections c
		         JOIN questions q ON q.collection_id = c.id
		         JOIN answers a ON q.id = a.question_id
		WHERE c.id = ?`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	fullCollection := &model.FullCollection{
		Collection: &model.Collection{},
		Questions:  []*model.FullQuestion{},
	}
	questionMap := make(map[string]*model.FullQuestion)
	for rows.Next() {
		var (
			cID            int
			cName          string
			cCreatedAt     time.Time
			cUpdatedAt     time.Time
			qID            int
			qQuestionField string
			qCreatedAt     time.Time
			aID            int
			aCreatedAt     time.Time
			aIsTrue        bool
			aAnswerField   string
		)
		if err := rows.Scan(&cID, &cName, &cCreatedAt, &cUpdatedAt, &qID, &qQuestionField, &qCreatedAt, &aID, &aCreatedAt, &aIsTrue, &aAnswerField); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		fullCollection.Collection = &model.Collection{
			ID:        fmt.Sprintf("%d", cID),
			Name:      cName,
			CreatedAt: cCreatedAt.Format(time.RFC3339),
			UpdatedAt: cUpdatedAt.Format(time.RFC3339),
		}
		question, exists := questionMap[fmt.Sprintf("%d", qID)]
		if !exists {
			question = &model.FullQuestion{
				ID:            fmt.Sprintf("%d", qID),
				QuestionField: qQuestionField,
				CreatedAt:     qCreatedAt.Format(time.RFC3339),
				Answers:       []*model.Answer{},
			}
			questionMap[fmt.Sprintf("%d", qID)] = question
			fullCollection.Questions = append(fullCollection.Questions, question)
		}
		question.Answers = append(question.Answers, &model.Answer{
			ID:          fmt.Sprintf("%d", aID),
			IsTrue:      aIsTrue,
			QuestionID:  fmt.Sprintf("%d", qID),
			AnswerField: aAnswerField,
			CreatedAt:   aCreatedAt.Format(time.RFC3339),
			UpdatedAt:   "",
		})
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return fullCollection, nil
}

// Update updates an existing collection in the database.
func (r *CollectionRepository) Update(collection *model.Collection) error {
	_, err := r.db.Exec("UPDATE collections SET name = ?, updated_at = ? WHERE id = ?",
		collection.Name, collection.UpdatedAt, collection.ID)
	if err != nil {
		log.Printf("Error updating collection: %v", err)
		return err
	}
	return nil
}

// Delete removes a collection from the database.
func (r *CollectionRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM collections WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting collection: %v", err)
		return err
	}
	return nil
}
