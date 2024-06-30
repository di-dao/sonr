package fs

import (
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

func CreateFSEntry(f Folder) *FSEntry {
	return &FSEntry{
		LocalPath: f.Path(),
		Address:   f.Name(),
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
