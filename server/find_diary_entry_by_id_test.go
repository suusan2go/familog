package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func TestFindDiaryEntryById(t *testing.T) {
	ctx := context.Background()
	req := &familog.FindDiaryEntryByIdRequest{}

	res, err := cli.FindDiaryEntryById(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
