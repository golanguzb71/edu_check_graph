package repository

import (
	"database/sql"
	"edu_test_graph/graph/model"
	utils "edu_test_graph/internal"
	database "edu_test_graph/internal/config"
	"errors"
	"fmt"
	"log"
)

type AnswerRepository struct {
	db *sql.DB
}

func NewAnswerRepository(db *sql.DB) *AnswerRepository {
	return &AnswerRepository{db: db}
}

func (r *AnswerRepository) Create(answer *model.Answer) error {
	_, err := r.db.Exec("INSERT INTO answers (is_true, question_id, answer_field) VALUES ($1, $2, $3)",
		answer.IsTrue, answer.QuestionID, answer.AnswerField)
	if err != nil {
		log.Printf("Error inserting answer: %v", err)
		return err
	}
	return nil
}

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

func (r *AnswerRepository) collectionExists(collectionID string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM collections WHERE id = $1)`
	err := r.db.QueryRow(query, collectionID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check if collection exists: %v", err)
	}
	return exists, nil
}

func (r *AnswerRepository) CreateInsertAnswer(answers model.AnswerInsert) error {
	key := answers.Key
	resp, err := utils.SearchByValue(database.RDB, key)
	if err != nil {
		return errors.New("not allowed; please insert a valid key: https://t.me/codevanbot")
	}
	userId := resp.UserID
	collectionID := answers.CollectionID
	fmt.Println(userId)
	fmt.Println(collectionID)

	exists, err := r.collectionExists(collectionID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("collection ID %d does not exist", collectionID)
	}

	var creatingModuleSql = `INSERT INTO student_collection(student_id, collection_id) VALUES ($1, $2) RETURNING id`
	var id int64
	err = r.db.QueryRow(creatingModuleSql, userId, collectionID).Scan(&id)
	if err != nil {
		return fmt.Errorf("failed to insert into student_collection: %v", err)
	}

	for _, el := range answers.Answers {
		checker := `SELECT EXISTS(SELECT 1 FROM answers WHERE id=$1 and question_id=$2)`
		var checking bool
		_ = r.db.QueryRow(checker, el.AnswerID, el.QuestionID).Scan(&checking)
		if !checking {
			return errors.New("not allowed answer_id not match with question_id")
		}
		err = r.db.QueryRow(`SELECT is_true FROM answers where id=$1`, el.AnswerID).Scan(&checker)
		if err != nil {
			return err
		}

		var creatingStudentAnswerSql = `INSERT INTO student_answer(student_collection_id, question_id, answer_id , is_true) VALUES ($1 , $2 , $3 , $4)`
		_, err = r.db.Exec(creatingStudentAnswerSql, id, el.QuestionID, el.AnswerID, checker)
		if err != nil {
			return err
		}
	}
	return nil
}
