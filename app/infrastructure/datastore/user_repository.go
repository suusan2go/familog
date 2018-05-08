package datastore

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/datastore"
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
		return fmt.Errorf("datastoredb: could not put User: %v", err)
	}
	u.ID = domain.UserID(k.ID)
	return nil
}
