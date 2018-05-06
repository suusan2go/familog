package usecase

import (
	"github.com/suusan2go/familog/app/domain"
)

type PublishDiaryEntryUsecase struct {
	diaryEntryRepo domain.DiaryEntryRepository
}

// PublishDiaryEntryInput input parameter
type PublishDiaryEntryInput struct {
	Body     string
	Images   []string
	AuthorID domain.UserID
	Emoji    string
}

// PublishDiaryEntryOutput output parameter
type PublishDiaryEntryOutput struct {
	DiaryEntry domain.DiaryEntry
}

func NewPublishDiaryEntryUsecase(diaryEntryRepo domain.DiaryEntryRepository) PublishDiaryEntryUsecase {
	return PublishDiaryEntryUsecase{diaryEntryRepo: diaryEntryRepo}
}

func (us *PublishDiaryEntryUsecase) PublishDiary(input PublishDiaryEntryInput) (*PublishDiaryEntryOutput, error) {
	de := &domain.DiaryEntry{
		AuthorID: input.AuthorID,
		Body:     input.Body,
		Images:   input.Images,
		Emoji:    input.Emoji,
	}
	if err := us.diaryEntryRepo.Save(de); err != nil {
		return nil, err
	}
	return &PublishDiaryEntryOutput{DiaryEntry: *de}, nil
}
