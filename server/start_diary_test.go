package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func TestStartDiary(t *testing.T) {
	ctx := context.Background()
	req := &familog.StartDiaryRequest{}

	res, err := cli.StartDiary(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
