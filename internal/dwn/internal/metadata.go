package internal

import (
	"github.com/di-dao/sonr/internal/orm"
	fs "github.com/di-dao/sonr/internal/vfs"
)

type Status interface{}

type info struct {
	Metadata *orm.Metadata
	File     fs.File
}

func InitMetadata(folder fs.Folder) (Status, error) {
	md := &orm.Metadata{}
	return &info{Metadata: md}, nil
}
