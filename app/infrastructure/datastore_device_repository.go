package infrastructure

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/suusan2go/familog/app/domain"
)

const deviceKind = "Device"

// DatastoreDeviceRepository implementation for DeviceRepository by Datastore
type DatastoreDeviceRepository struct {
	client *datastore.Client
}

// NewDatastoreDeviceRepository returns NewDatastoreDeviceRepository struct
func NewDatastoreDeviceRepository(c *datastore.Client) DatastoreDeviceRepository {
	return DatastoreDeviceRepository{client: c}
}

// Save save Device to Datastore
func (rp DatastoreDeviceRepository) Save(d *domain.Device) error {
	ctx := context.Background()
	var k *datastore.Key
	if d.ID == 0 {
		k = datastore.IncompleteKey(deviceKind, nil)
	} else {
		k = datastore.IDKey(deviceKind, int64(d.ID), nil)
	}
	k, err := rp.client.Put(ctx, k, d)
	if err != nil {
		return fmt.Errorf("datastoredb: could not put Book: %v", err)
	}
	d.ID = uint64(k.ID)
	return nil
}

// FindByDeviceID find device from datastore
func (rp DatastoreDeviceRepository) FindByDeviceID(deviceID string) (*domain.Device, error) {
	ctx := context.Background()
	q := datastore.NewQuery(deviceKind).Filter("DeviceId =", deviceID).Limit(1)
	devices := make([]*domain.Device, 0)

	keys, err := rp.client.GetAll(ctx, q, &devices)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list devices: %v", err)
	}

	for i, k := range keys {
		devices[i].ID = uint64(k.ID)
	}

	return devices[0], nil
}
