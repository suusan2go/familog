package domain

import (
	"time"
)

// DiaryEntryID value object
type DiaryEntryID int64

// DiaryEntry struct
type DiaryEntry struct {
	ID        DiaryEntryID
	AuthorID  UserID
	Body      string
	Emoji     string
	Images    []string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (d *DiaryEntry) PrimaryImage() string {
	if len(d.Images) == 0 {
		return ""
	}
	return d.Images[0]
}
