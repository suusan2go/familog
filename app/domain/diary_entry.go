package domain

import (
	"time"
)

// DiaryEntryID value object
type DiaryEntryID int64

// DiaryEntry struct
type DiaryEntry struct {
	ID        DiaryEntryID
	Title     string
	Body      string
	Emoji     string
	Images    []string
	UpdatedAt time.Time
	CreatedAt time.Time
}
