package server

import (
	"errors"

	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func (s FamilogServer) FindDiaryEntries(ctx context.Context, r *familog.FindDiaryEntriesRequest) (*familog.FindDiaryEntriesResponse, error) {
	return nil, errors.New("not yet implemented")
}
