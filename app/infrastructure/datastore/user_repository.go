package datastore

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
	"github.com/suusan2go/familog/app/domain"
)

const userKind = "User"

// UserRepository implementation for domain.UserRepository by Datastore
type UserRepository struct {
	client *datastore.Client
}

// NewUserRepository returns NewUserRepository struct
func NewUserRepository(c *datastore.Client) UserRepository {
	return UserRepository{client: c}
}

// Save save domain.User to datastore
func (rp UserRepository) Save(u *domain.User) error {
	ctx := context.Background()

	k := datastore.NameKey(userKind, string(u.ID), nil)
	err := rp.client.Get(ctx, k, &domain.User{}) // 対象のIDで存在するかの確認
	if err != nil {
		u.UpdatedAt = time.Now()
		u.CreatedAt = time.Now()
	} else {
		u.UpdatedAt = time.Now()
	}
	k, err = rp.client.Put(ctx, k, u)
	if err != nil {
		return mapError(errors.Wrap(err, "datastoredb: could not put User"))
	}
	u.ID = domain.UserID(k.ID)
	return nil
}
