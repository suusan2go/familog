package server

import (
	"errors"

	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func (s FamilogServer) PublishDiaryEntry(ctx context.Context, r *familog.PublishDiaryEntryRequest) (*familog.PublishDiaryEntryResponse, error) {
	return nil, errors.New("not yet implemented")
}
