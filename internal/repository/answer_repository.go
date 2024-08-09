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

func (r *AnswerRepository) CreateInsertAnswer(answers model.AnswerInsert) (*model.CommonResponse, error) {
	key := answers.Key
	resp, err := utils.SearchByValue(database.RDB, key)
	if err != nil {
		return &model.CommonResponse{}, errors.New("not allowed; please insert a valid key: https://t.me/codevanbot")
	}
	userId := resp.UserID
	collectionID := answers.CollectionID
	fmt.Println(userId)
	fmt.Println(collectionID)

	exists, err := r.collectionExists(collectionID)
	if err != nil {
		return &model.CommonResponse{}, err
	}
	if !exists {
		return &model.CommonResponse{}, fmt.Errorf("collection ID %d does not exist", collectionID)
	}

	var creatingModuleSql = `INSERT INTO student_collection(student_id, collection_id) VALUES ($1, $2) RETURNING id`
	var id int64
	err = r.db.QueryRow(creatingModuleSql, userId, collectionID).Scan(&id)
	if err != nil {
		return &model.CommonResponse{}, fmt.Errorf("failed to insert into student_collection: %v", err)
	}

	var result []*model.ResponseAfterTesting
	var TAC = 0
	for _, el := range answers.Answers {
		checker := `SELECT EXISTS(SELECT 1 FROM answers WHERE id=$1 and question_id=$2)`
		var checking bool
		_ = r.db.QueryRow(checker, el.AnswerID, el.QuestionID).Scan(&checking)
		if !checking {
			return &model.CommonResponse{},
				errors.New("not allowed answer_id not match with question_id")
		}
		err = r.db.QueryRow(`SELECT is_true FROM answers where id=$1`, el.AnswerID).Scan(&checking)
		if err != nil {
			return &model.CommonResponse{}, err
		}
		if checking {
			TAC += 1
		}
		var PR model.ResponseAfterTesting
		PR.IsTrue = checking
		_ = r.db.QueryRow("SELECT question_field FROM questions where id=$1", el.QuestionID).Scan(&PR.QuestionField)
		_ = r.db.QueryRow("SELECT answer_field FROM answers where id=$1", el.AnswerID).Scan(&PR.GivenAnswerField)
		_ = r.db.QueryRow("SELECT answer_field FROM answers where is_true=true and question_id=$1", el.QuestionID).Scan(&PR.TrueAnswerField)
		var creatingStudentAnswerSql = `INSERT INTO student_answer(student_collection_id, question_id, answer_id , is_true) VALUES ($1 , $2 , $3 , $4)`
		_, err = r.db.Exec(creatingStudentAnswerSql, id, el.QuestionID, el.AnswerID, checking)
		if err != nil {
			return &model.CommonResponse{}, err
		}
		result = append(result, &PR)
	}
	var CQC int
	var query = `SELECT COUNT(id)
	FROM questions
	where collection_id = $1`
	_ = r.db.QueryRow(query, collectionID).Scan(&CQC)
	return makingResponse(result, CQC, TAC)
}

func makingResponse(result []*model.ResponseAfterTesting, commonQuestionCount, trueAnswerCount int) (*model.CommonResponse, error) {
	res := &model.CommonResponse{}
	res.Answers = result
	pureProportion := trueAnswerCount * 100 / commonQuestionCount
	var level string

	switch {
	case pureProportion >= 0 && pureProportion <= 19:
		level = "BEGINNER"
	case pureProportion >= 20 && pureProportion <= 39:
		level = "ELEMENTARY"
	case pureProportion >= 40 && pureProportion <= 59:
		level = "PRE_INTERMEDIATE"
	case pureProportion >= 60 && pureProportion <= 74:
		level = "INTERMEDIATE"
	case pureProportion >= 75 && pureProportion <= 89:
		level = "UPPER_INTERMEDIATE"
	case pureProportion >= 90 && pureProportion <= 99:
		level = "ADVANCED"
	case pureProportion == 100:
		level = "PROFICIENT"
	default:
		level = "BEGINNER"
	}

	res.Message = fmt.Sprintf("Siz %d ta savoldan %d tasiga tog'ri javob berdingiz va sizning darajangiz %s deb belgilandi. Sizga Taklif qilinadigan guruhlarimiz hamda ularning o'qituvchilari bilan tanishing", commonQuestionCount, trueAnswerCount, level)
	rows, _ := database.DB.Query(`SELECT id, name, teacher_name, level, created_at, updated_at FROM groups where level=$1`, level)
	var groups []*model.Group
	for rows.Next() {
		var group model.Group
		_ = rows.Scan(&group.ID, &group.Name, &group.TeacherName, &group.Level, &group.CreatedAt, &group.UpdatedAt)
		groups = append(groups, &group)
	}

	res.RequestGroup = groups
	return res, nil
}
