package app

import (
	"cloud.google.com/go/datastore"
	"github.com/suusan2go/familog/app/domain"
	infrads "github.com/suusan2go/familog/app/infrastructure/datastore"
)

// Registry service registry for Dependency Injection
type Registry struct {
	dsc *datastore.Client
}

func NewRegistry(c *datastore.Client) *Registry {
	return &Registry{dsc: c}
}

// DiaryRepository returns implemented Repository
func (r Registry) DiaryRepository() domain.DiaryRepository {
	return infrads.NewDatastoreDiaryRepository(r.dsc)
}

//// UserRepository returns implemented Struct
//func (r Registry) UserRepository() domain.UserRepository {
//	return infrastructure.NewDatastoreUserRepository(r.dsc)
//}
