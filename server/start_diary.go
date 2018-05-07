package server

import (
	"github.com/suusan2go/familog"
	"github.com/suusan2go/familog/app/domain"
	"github.com/suusan2go/familog/app/usecase"
	context "golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s FamilogServer) StartDiary(ctx context.Context, r *familog.StartDiaryRequest) (*familog.StartDiaryResponse, error) {
	auth, err := s.currentAuth(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Failed to authenticate user "+err.Error())
	}
	us := usecase.NewStartDiaryUsecase(s.Registry.DiaryRepository())
	out, err := us.StartDiary(usecase.StartDiaryInput{Title: r.Title, UserID: domain.UserID(auth.ID)})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to start diary")
	}
	d := familog.Diary{
		Id:    int64(out.Diary.ID),
		Title: out.Diary.Title,
	}
	return &familog.StartDiaryResponse{Diary: &d}, nil
}
