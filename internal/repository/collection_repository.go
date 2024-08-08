package repository

import (
	"database/sql"
	"edu_test_graph/graph/model"
	"fmt"
	"log"
	"time"
)

type CollectionRepository struct {
	db *sql.DB
}

func NewCollectionRepository(db *sql.DB) *CollectionRepository {
	return &CollectionRepository{db: db}
}

// Create inserts a new collection into the database.
func (r *CollectionRepository) Create(collection *model.Collection) error {
	_, err := r.db.Exec("INSERT INTO collections (name) VALUES ($1)", collection.Name)
	if err != nil {
		log.Printf("Error inserting collection: %v", err)
		return err
	}
	return nil
}

func (r *CollectionRepository) Get(id int) (*model.FullCollection, error) {
	rows, err := r.db.Query(`
		SELECT c.id,
		       c.name,
		       c.created_at,
		       c.updated_at,
		       q.id,
		       q.question_field,
		       q.created_at,
		       COALESCE(a.id, 0) AS answer_id,
		       COALESCE(a.created_at, '1970-01-01T00:00:00Z') AS answer_created_at,
		       COALESCE(a.is_true, false) AS is_true,
		       COALESCE(a.answer_field, '') AS answer_field
		FROM collections c
		JOIN questions q ON q.collection_id = c.id
		LEFT JOIN answers a ON q.id = a.question_id
		WHERE c.id = $1`, id)
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

		// Set collection details only once
		if fullCollection.Collection.ID == "" {
			fullCollection.Collection = &model.Collection{
				ID:        fmt.Sprintf("%d", cID),
				Name:      cName,
				CreatedAt: cCreatedAt.Format(time.RFC3339),
				UpdatedAt: cUpdatedAt.Format(time.RFC3339),
			}
		}

		// Add questions and answers
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

		// Only add the answer if it exists
		if aID != 0 {
			question.Answers = append(question.Answers, &model.Answer{
				ID:          fmt.Sprintf("%d", aID),
				IsTrue:      aIsTrue,
				QuestionID:  fmt.Sprintf("%d", qID),
				AnswerField: aAnswerField,
				CreatedAt:   aCreatedAt.Format(time.RFC3339),
				UpdatedAt:   "", // Set to empty if no update timestamp available
			})
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return fullCollection, nil
}

func (r *CollectionRepository) Update(collection *model.Collection) error {
	_, err := r.db.Exec("UPDATE collections SET name = $1, updated_at = $2 WHERE id = $3",
		collection.Name, time.Now(), collection.ID)
	if err != nil {
		log.Printf("Error updating collection: %v", err)
		return err
	}
	return nil
}

func (r *CollectionRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM collections WHERE id = $1", id)
	if err != nil {
		log.Printf("Error deleting collection: %v", err)
		return err
	}
	return nil
}

func (r *CollectionRepository) GetCollections() ([]*model.Collection, error) {
	rows, err := r.db.Query(`SELECT id, name, created_at, updated_at FROM collections`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var collections []*model.Collection
	for rows.Next() {
		var collect model.Collection
		err = rows.Scan(&collect.ID, &collect.Name, &collect.CreatedAt, &collect.UpdatedAt)
		if err != nil {
			return nil, err
		}
		collections = append(collections, &collect)
	}
	return collections, nil
}
