package vault

import (
	"encoding/json"
	"fmt"

	"go.etcd.io/bbolt"

	"github.com/IkramCode/vault/internals/crypto"
)

type VaultEntry struct {
	Site     string `json:"site"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func AddEntry(db *VaultDB, site, username, password, masterPassword string) error {
	salt, _ := crypto.GenerateSalt()
	key, err := crypto.MasterKey(masterPassword, salt)
	if err != nil {
		return err
	}
	entry := VaultEntry{Site: site, Username: username, Password: password}
	blob, _ := json.Marshal(entry)
	encrypted, err := crypto.Encrypt(key, blob)
	if err != nil {
		return err
	}
	data := append(salt, encrypted...)
	return db.Put(site, data)
}

func GetEntry(db *VaultDB, site, masterPass string) (*VaultEntry, error) {
	data, err := db.Get(site)
	if err != nil {
		return nil, err
	}
	saltsize := 16
	if len(data) < saltsize {
		return nil, fmt.Errorf("stored data too short")
	}
	salt, encrypted := data[:saltsize], data[saltsize:]
	key, err := crypto.MasterKey(masterPass, salt)
	if err != nil {
		return nil, err
	}
	blob, err := crypto.Decrypt(key, encrypted)
	if err != nil {
		return nil, err
	}
	var entry VaultEntry
	if err := json.Unmarshal(blob, &entry); err != nil {
		return nil, err
	}
	return &entry, nil
}

func ListEntries(db *VaultDB) ([]string, error) {
	var sites []string
	err := db.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			sites = append(sites, string(k))
		}
		return nil
	})
	return sites, err
}
