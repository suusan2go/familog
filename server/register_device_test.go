package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suusan2go/familog"
	context "golang.org/x/net/context"
)

func TestRegisterDevice(t *testing.T) {
	ctx := context.Background()
	req := &familog.RegisterDeviceRequest{}

	res, err := cli.RegisterDevice(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
