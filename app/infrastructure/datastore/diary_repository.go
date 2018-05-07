package datastore

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/suusan2go/familog/app/domain"
)

// DiaryRepository implements domain.DiaryRepository by Datastore
type DiaryRepository struct {
	client *datastore.Client
}

const diaryKind = "Diary"

func (repo *DiaryRepository) FindByID(id domain.DiaryID) (*domain.Diary, error) {
	ctx := context.Background()
	key := datastore.IDKey(diaryKind, int64(id), nil)
	fmt.Printf("%v", key)
	diary := &domain.Diary{}
	if err := repo.client.Get(ctx, key, diary); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get Diary: %v", err)
	}
	diary.ID = id
	return diary, nil
}

func NewDatastoreDiaryRepository(client *datastore.Client) *DiaryRepository {
	return &DiaryRepository{client}
}

func (repo *DiaryRepository) Save(d *domain.Diary) error {
	ctx := context.Background()
	var k *datastore.Key
	if d.ID == 0 {
		k = datastore.IncompleteKey(diaryKind, nil)
		d.UpdatedAt = time.Now()
		d.CreatedAt = time.Now()
	} else {
		k = datastore.IDKey(diaryKind, int64(d.ID), nil)
		d.UpdatedAt = time.Now()
	}
	k, err := repo.client.Put(ctx, k, d)
	if err != nil {
		return fmt.Errorf("datastoredb: could not put Diary: %v", err)
	}
	d.ID = domain.DiaryID(k.ID)
	return nil
}

func (repo *DiaryRepository) FindSubscribes(u *domain.UserID) ([]domain.Diary, error) {
	panic("implement me")
}
