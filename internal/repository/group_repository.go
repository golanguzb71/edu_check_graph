package repository

import (
	"database/sql"
	"edu_test_graph/graph/model"
	"fmt"
	"log"
	"time"
)

type GroupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

// Create inserts a new group into the database.
func (r *GroupRepository) Create(group *model.Group) error {
	_, err := r.db.Exec("INSERT INTO groups (name, teacher_name, level) VALUES ($1, $2, $3)",
		group.Name, group.TeacherName, group.Level)
	if err != nil {
		log.Printf("Error inserting group: %v", err)
		return err
	}
	return nil
}

func (r *GroupRepository) Get(id *string, orderLevel *bool) ([]*model.Group, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if id != nil {
		sql := `SELECT id, name, teacher_name, level, created_at, updated_at FROM groups WHERE id = $1`
		rows, err = r.db.Query(sql, id)
	} else {
		sql := `SELECT id, name, teacher_name, level, created_at, updated_at FROM groups`
		if orderLevel != nil {
			sql += ` order by level`
		}
		rows, err = r.db.Query(sql)
	}

	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var groups []*model.Group

	for rows.Next() {
		var (
			id          int
			name        string
			teacherName string
			level       string
			createdAt   time.Time
			updatedAt   time.Time
		)

		if err := rows.Scan(&id, &name, &teacherName, &level, &createdAt, &updatedAt); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		group := &model.Group{
			ID:          fmt.Sprintf("%d", id),
			Name:        name,
			TeacherName: teacherName,
			Level:       level,
			CreatedAt:   createdAt.Format(time.RFC3339),
			UpdatedAt:   updatedAt.Format(time.RFC3339),
		}

		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return groups, nil
}

func (r *GroupRepository) Update(group *model.Group) error {
	_, err := r.db.Exec("UPDATE groups SET name = $1, teacher_name = $2, level = $3, updated_at = $4 WHERE id = $5",
		group.Name, group.TeacherName, group.Level, time.Now(), group.ID)
	if err != nil {
		log.Printf("Error updating group: %v", err)
		return err
	}
	return nil
}

// Delete removes a group from the database.
func (r *GroupRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM groups WHERE id = $1", id)
	if err != nil {
		log.Printf("Error deleting group: %v", err)
		return err
	}
	return nil
}
