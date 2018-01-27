package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func TestPublishDiaryEntry(t *testing.T) {
	ctx := context.Background()
	req := &familog.PublishDiaryEntryRequest{}

	res, err := cli.PublishDiaryEntry(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
