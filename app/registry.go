package app

import (
	"cloud.google.com/go/datastore"
	"github.com/suusan2go/familog/app/domain"
	"github.com/suusan2go/familog/app/infrastructure"
)

// Registry service registry for Dependency Injection
type Registry struct {
	dsc *datastore.Client
}

func NewRegistry(c *datastore.Client) Registry {
	return Registry{dsc: c}
}

// DeviceRepository returns implemented Struct
func (r Registry) DeviceRepository() domain.DeviceRepository {
	return infrastructure.NewDatastoreDeviceRepository(r.dsc)
}

// UserRepository returns implemented Struct
func (r Registry) UserRepository() domain.UserRepository {
	return infrastructure.NewDatastoreUserRepository(r.dsc)
}
