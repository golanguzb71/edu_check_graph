package repository

import (
	"database/sql"
	"edu_test_graph/graph/model"
	"log"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) Create(student *model.Student) error {
	_, err := r.db.Exec("INSERT INTO students (phone_number, full_name) VALUES (?, ?)",
		student.PhoneNumber, student.FullName)
	if err != nil {
		log.Printf("Error inserting student: %v", err)
		return err
	}
	return nil
}

func (r *StudentRepository) Get(id int) (*model.Student, error) {
	row := r.db.QueryRow("SELECT id, phone_number, full_name, created_at, updated_at FROM students WHERE id = ?", id)
	student := &model.Student{}
	err := row.Scan(&student.ID, &student.PhoneNumber, &student.FullName, &student.CreatedAt, &student.UpdatedAt)
	if err != nil {
		log.Printf("Error retrieving student: %v", err)
		return nil, err
	}
	return student, nil
}

func (r *StudentRepository) Update(student *model.Student) error {
	_, err := r.db.Exec("UPDATE students SET phone_number = ?, full_name = ?, updated_at = ? WHERE id = ?",
		student.PhoneNumber, student.FullName, student.UpdatedAt, student.ID)
	if err != nil {
		log.Printf("Error updating student: %v", err)
		return err
	}
	return nil
}

func (r *StudentRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM students WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting student: %v", err)
		return err
	}
	return nil
}
