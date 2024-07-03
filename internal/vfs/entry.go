package fs

import (
	"time"

	"gorm.io/gorm"
)

type FSEntry struct {
	gorm.Model
	LocalPath       string
	IPFSPath        string
	Address         string
	Synced          bool
	Encrypted       bool
	Published       bool
	LastSyncTime    time.Time
	LastPublishTime time.Time
}

func CreateFSEntry(f Folder) *FSEntry {
	return &FSEntry{
		LocalPath:       f.Path(),
		Address:         f.Name(),
		Synced:          false,
		Encrypted:       false,
		Published:       false,
		LastSyncTime:    time.Now(),
		LastPublishTime: time.Now(),
	}
}

func (f *FSEntry) SetEncrypted() {
	f.Encrypted = true
}

func (f *FSEntry) SetUnencrypted() {
	f.Encrypted = false
}

func (f *FSEntry) SetSynced(ipfsPath string) {
	f.Synced = true
	f.IPFSPath = ipfsPath
	f.LastSyncTime = time.Now()
}

func (f *FSEntry) SetPublished() {
	f.Published = true
	f.LastPublishTime = time.Now()
}
