package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func TestUpdateProfile(t *testing.T) {
	ctx := context.Background()
	req := &familog.UpdateProfileRequest{}

	res, err := cli.UpdateProfile(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
