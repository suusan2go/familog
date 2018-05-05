package domain

type DiaryEntryRepository interface {
	Save(d *DiaryEntry) error
	FindSubscribes(u *UserID) ([]DiaryEntry, error)
}
