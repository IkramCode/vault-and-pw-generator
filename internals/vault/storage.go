package vault

import (
	"fmt"
	"go.etcd.io/bbolt"
)

type VaultDB struct {
	db *bbolt.DB
}

const bucketName = "entries"

func Open(path string) (*VaultDB, error) {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	if err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	}); err != nil {
		return nil, err
	}
	return &VaultDB{db: db}, nil
}

func (v *VaultDB) Put(key string, value []byte) error {
	return v.db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucketName)).Put([]byte(key), value)
	})
}

func (v *VaultDB) Get(key string) ([]byte, error) {
	var value []byte
	err := v.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		value = b.Get([]byte(key))
		return nil
	})
	if err != nil {
		return nil, err
	}
	if value == nil {
		return nil, fmt.Errorf("entry not found: %s", key)
	}
	return value, nil
}

func (v *VaultDB) Close() error { return v.db.Close() }
