package mysql

import (
	"database/sql"
	"errors"

	"satotory.dyplom.go/pkg/models"
)

type LessonModel struct {
	DB *sql.DB
}

func (m *LessonModel) Get(id int) (*models.Lesson, error) {
	s := &models.Lesson{}
	err := m.DB.QueryRow("SELECT id, title, content, created FROM lessons WHERE id = ?", id).Scan(&s.ID, &s.Title, &s.Content, &s.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *LessonModel) Latest() ([]*models.Lesson, error) {
	return nil, nil
}
