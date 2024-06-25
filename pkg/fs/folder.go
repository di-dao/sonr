package fs

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/ipfs/boxo/files"
)

type Folder string

func (f Folder) Path() string {
	return string(f)
}

func (f Folder) Node() files.Node {
	entries, err := f.ReadDir()
	if err != nil {
		return nil
	}

	fileList := make([]files.DirEntry, 0, len(entries))
	for _, entry := range entries {
		name := entry.Name()
		path := filepath.Join(f.Path(), name)
		if entry.IsDir() {
			fileList = append(fileList, files.FileEntry(name, files.NewSerialFile(name, path, false)))
		} else {
			fileList = append(fileList, files.FileEntry(name, files.NewSerialFile(name, path, true)))
		}
	}

	return files.NewSliceDirectory(fileList)
}

// Create creates the folder if it doesn't exist
func (f Folder) Create() error {
	return os.MkdirAll(f.Path(), os.ModePerm)
}

// Exists checks if the folder exists
func (f Folder) Exists() bool {
	_, err := os.Stat(f.Path())
	return !os.IsNotExist(err)
}

// Remove removes the folder and its contents
func (f Folder) Remove() error {
	return os.RemoveAll(f.Path())
}

// ReadDir reads the contents of the folder
func (f Folder) ReadDir() ([]fs.DirEntry, error) {
	return os.ReadDir(f.Path())
}

// Join joins the folder path with the given elements
func (f Folder) Join(elem ...string) Folder {
	return Folder(filepath.Join(append([]string{f.Path()}, elem...)...))
}

// IsDir checks if the folder is a directory
func (f Folder) IsDir() bool {
	info, err := os.Stat(f.Path())
	return err == nil && info.IsDir()
}
