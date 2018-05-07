package server

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/suusan2go/familog"
	"github.com/suusan2go/familog/app/domain"
	"github.com/suusan2go/familog/app/usecase"
	context "golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s FamilogServer) PublishDiaryEntry(ctx context.Context, r *familog.PublishDiaryEntryRequest) (*familog.PublishDiaryEntryResponse, error) {
	auth, err := s.currentAuth(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Failed to authenticate user "+err.Error())
	}
	us := usecase.NewPublishDiaryEntryUsecase(s.Registry.DiaryEntryRepository(), s.Registry.DiaryRepository())
	output, err := us.PublishDiaryEntry(usecase.PublishDiaryEntryInput{
		DiaryID:  domain.DiaryID(r.DiaryId),
		Body:     r.DiaryEntryForm.Body,
		Emoji:    r.DiaryEntryForm.Emoji,
		AuthorID: domain.UserID(auth.ID), // TODO
		Images:   r.DiaryEntryForm.ImageUrls,
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to publish Diary"+err.Error())
	}
	return &familog.PublishDiaryEntryResponse{
		DiaryEntry: &familog.DiaryEntry{
			Id:        int64(output.DiaryEntry.ID),
			Body:      output.DiaryEntry.Body,
			Emoji:     output.DiaryEntry.Emoji,
			ImageUrls: output.DiaryEntry.Images,
			PublishedAt: &timestamp.Timestamp{
				Seconds: int64(output.DiaryEntry.CreatedAt.Second()),
				Nanos:   int32(output.DiaryEntry.CreatedAt.Nanosecond()),
			},
		},
	}, nil
}
