package internal

import (
	"github.com/di-dao/sonr/internal/models"
	"github.com/di-dao/sonr/pkg/fs"
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
