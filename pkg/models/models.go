package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: record not found")

type Lesson struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}
