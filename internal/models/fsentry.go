package models

import (
	"gorm.io/gorm"
)

type FSEntry struct {
	gorm.Model
	LocalPath string
	IPFSPath  string
	Address   string
}
