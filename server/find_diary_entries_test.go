package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func TestFindDiaryEntries(t *testing.T) {
	ctx := context.Background()
	req := &familog.FindDiaryEntriesRequest{}

	res, err := cli.FindDiaryEntries(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
