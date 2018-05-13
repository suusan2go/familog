package datastore

import (
	"cloud.google.com/go/datastore"
	"github.com/suusan2go/familog/app/domain"
)

func mapError(e error) error {
	if e == datastore.ErrNoSuchEntity {
		return &domain.EntityNotFoundError{e}
	}
	return e
}
