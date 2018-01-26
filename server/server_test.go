package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/lileio/lile"
	"github.com/suusan2go/familog"
)

var s = FamilogServer{}
var cli familog.FamilogClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		familog.RegisterFamilogServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = familog.NewFamilogClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
