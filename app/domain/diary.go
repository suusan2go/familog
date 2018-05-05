package domain

import (
	"time"
)

// DiaryID value object
type DiaryID int64

// Diary struct
type Diary struct {
	ID          DiaryID
	Title       string
	Subscribers []UserID
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

func (d *Diary) AddSubscriber(id UserID) {
	d.Subscribers = append(d.Subscribers, id)
}

// NewDiary create new diary struct
func NewDiary(title string) *Diary {
	current := time.Now()
	return &Diary{
		Title:     title,
		UpdatedAt: current,
		CreatedAt: current,
	}
}
