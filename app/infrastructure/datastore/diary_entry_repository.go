package datastore

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
	"github.com/suusan2go/familog/app/domain"
)

// DiaryEntryRepository implements domain.DiaryRepository by Datastore
type DiaryEntryRepository struct {
	client *datastore.Client
}

var diaryEntryKind = "DiaryEntry"

func NewDatastoreDiaryEntryRepository(client *datastore.Client) *DiaryEntryRepository {
	return &DiaryEntryRepository{client}
}

func (repo *DiaryEntryRepository) Save(d *domain.DiaryEntry) error {
	ctx := context.Background()
	var k *datastore.Key
	if d.ID == 0 {
		k = datastore.IncompleteKey(diaryEntryKind, nil)
		d.UpdatedAt = time.Now()
		d.CreatedAt = time.Now()
	} else {
		k = datastore.IDKey(diaryEntryKind, int64(d.ID), nil)
		d.UpdatedAt = time.Now()
	}
	k, err := repo.client.Put(ctx, k, d)
	if err != nil {
		return mapError(errors.Wrap(err, "datastoredb: could not put DiaryEntry"))
	}
	d.ID = domain.DiaryEntryID(k.ID)
	return nil
}

func (repo *DiaryEntryRepository) FindSubscribes(u *domain.UserID) ([]domain.DiaryEntry, error) {
	panic("implement me")
}
