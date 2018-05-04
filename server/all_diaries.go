package server

import (
	// "errors"

	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func (s FamilogServer) AllDiaries(ctx context.Context, r *familog.AllDiariesRequest) (*familog.AllDiariesResponse, error) {
	res := familog.AllDiariesResponse{}
	res.Diaries = []*familog.Diary{}
	return &res, nil
}
