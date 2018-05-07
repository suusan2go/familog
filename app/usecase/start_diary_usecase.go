package usecase

import (
	"github.com/suusan2go/familog/app/domain"
)

// StartDiaryUsecase usecase struct
type StartDiaryUsecase struct {
	diaryRepo domain.DiaryRepository
}

// StartDiaryInput input parameter
type StartDiaryInput struct {
	Title  string
	UserID domain.UserID
}

// StartDiaryOutput output parameter
type StartDiaryOutput struct {
	Diary domain.Diary
}

func NewStartDiaryUsecase(diaryRepo domain.DiaryRepository) StartDiaryUsecase {
	return StartDiaryUsecase{diaryRepo: diaryRepo}
}

// StartDiary start new diary
func (us *StartDiaryUsecase) StartDiary(input StartDiaryInput) (*StartDiaryOutput, error) {
	diary := domain.NewDiary(input.Title)
	diary.AddSubscriber(input.UserID)
	if err := us.diaryRepo.Save(diary); err != nil {
		return nil, err
	}
	return &StartDiaryOutput{Diary: *diary}, nil
}
