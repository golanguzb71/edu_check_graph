package repository

import (
	"database/sql"
	"edu_test_graph/graph/model"
	"log"
)

type GroupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

// Create inserts a new group into the database.
func (r *GroupRepository) Create(group *model.Group) error {
	_, err := r.db.Exec("INSERT INTO groups (name, teacher_name, level) VALUES (?, ?, ?)",
		group.Name, group.TeacherName, group.Level)
	if err != nil {
		log.Printf("Error inserting group: %v", err)
		return err
	}
	return nil
}

// Get retrieves a group by ID.
func (r *GroupRepository) Get(id int) (*model.Group, error) {
	row := r.db.QueryRow("SELECT id, name, teacher_name, level, created_at, updated_at FROM groups WHERE id = ?", id)
	group := &model.Group{}
	err := row.Scan(&group.ID, &group.Name, &group.TeacherName, &group.Level, &group.CreatedAt, &group.UpdatedAt)
	if err != nil {
		log.Printf("Error retrieving group: %v", err)
		return nil, err
	}
	return group, nil
}

// Update updates an existing group in the database.
func (r *GroupRepository) Update(group *model.Group) error {
	_, err := r.db.Exec("UPDATE groups SET name = ?, teacher_name = ?, level = ?, updated_at = ? WHERE id = ?",
		group.Name, group.TeacherName, group.Level, group.UpdatedAt, group.ID)
	if err != nil {
		log.Printf("Error updating group: %v", err)
		return err
	}
	return nil
}

// Delete removes a group from the database.
func (r *GroupRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM groups WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting group: %v", err)
		return err
	}
	return nil
}
