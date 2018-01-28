package server

import (
	"github.com/suusan2go/familog"
	"github.com/suusan2go/familog/app"
)

type FamilogServer struct {
	familog.FamilogServer
	registry app.Registry
}
