package usecase

import (
	"github.com/suusan2go/familog/app/domain"
)

type PublishDiaryEntryUsecase struct {
	diaryRepo      domain.DiaryRepository
	diaryEntryRepo domain.DiaryEntryRepository
}

// PublishDiaryEntryInput input parameter
type PublishDiaryEntryInput struct {
	Body     string
	Images   []string
	AuthorID domain.UserID
	Emoji    string
	DiaryID  domain.DiaryID
}

// PublishDiaryEntryOutput output parameter
type PublishDiaryEntryOutput struct {
	DiaryEntry domain.DiaryEntry
}

func NewPublishDiaryEntryUsecase(diaryEntryRepo domain.DiaryEntryRepository, diaryRepo domain.DiaryRepository) PublishDiaryEntryUsecase {
	return PublishDiaryEntryUsecase{diaryEntryRepo: diaryEntryRepo, diaryRepo: diaryRepo}
}

func (us *PublishDiaryEntryUsecase) PublishDiaryEntry(input PublishDiaryEntryInput) (*PublishDiaryEntryOutput, error) {
	d, err := us.diaryRepo.FindByID(input.DiaryID)
	if err != nil {
		return nil, err
	}
	de := d.AddEntry(
		input.AuthorID,
		input.Body,
		input.Images,
		input.Emoji,
	)
	if err := us.diaryEntryRepo.Save(de); err != nil {
		return nil, err
	}
	return &PublishDiaryEntryOutput{DiaryEntry: *de}, nil
}
