package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func TestSubscribeDiary(t *testing.T) {
	ctx := context.Background()
	req := &familog.SubscribeDiaryRequest{}

	res, err := cli.SubscribeDiary(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
