package domain

type DiaryRepository interface {
	Save(d *Diary) error
	FindSubscribes(u *UserID) ([]Diary, error)
	FindByID(id DiaryID) (Diary, error)
}
