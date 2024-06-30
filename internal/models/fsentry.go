package models

import (
	"path/filepath"

	"gorm.io/gorm"
)

type FSEntry struct {
	gorm.Model
	LocalPath string
	IPFSPath  string
	Address   string
	Synced    bool
	Encrypted bool
}

func CreateFSEntry(localPath string) *FSEntry {
	return &FSEntry{
		LocalPath: localPath,
		Address:   filepath.Base(localPath),
		Synced:    false,
		Encrypted: false,
	}
}

func (f *FSEntry) SetSynced() {
	f.Synced = true
}

func (f *FSEntry) SetEncrypted() {
	f.Encrypted = true
}

func (f *FSEntry) SetUnencrypted() {
	f.Encrypted = false
}
