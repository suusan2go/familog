package infrastructure

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/suusan2go/familog/app/domain"
)

const userKind = "User"

// DatastoreUserRepository implementation for domain.UserRepository by Datastore
type DatastoreUserRepository struct {
	client *datastore.Client
}

// NewDatastoreUserRepository returns NewDatastoreUserRepository struct
func NewDatastoreUserRepository(c *datastore.Client) DatastoreUserRepository {
	return DatastoreUserRepository{client: c}
}

// Save save doman.User to datastore
func (rp DatastoreUserRepository) Save(u *domain.User) error {
	ctx := context.Background()
	var k *datastore.Key
	if u.ID == 0 {
		k = datastore.IncompleteKey(userKind, nil)
		u.UpdatedAt = time.Now()
		u.CreatedAt = time.Now()
	} else {
		k = datastore.IDKey(userKind, int64(u.ID), nil)
		u.UpdatedAt = time.Now()
	}
	k, err := rp.client.Put(ctx, k, u)
	if err != nil {
		return fmt.Errorf("datastoredb: could not put User: %v", err)
	}
	u.ID = domain.UserID(k.ID)
	return nil
}
