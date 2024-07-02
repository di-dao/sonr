package internal

import (
	"github.com/di-dao/sonr/internal/fs"
	"github.com/di-dao/sonr/internal/models"
)

type Status interface{}

type info struct {
	Metadata *models.Metadata
	File     fs.File
}

func InitMetadata(folder fs.Folder) (Status, error) {
	md := &models.Metadata{}
	return &info{Metadata: md}, nil
}
