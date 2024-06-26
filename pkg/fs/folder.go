package fs

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// Folder represents a folder in the filesystem
type Folder string

// NewFolder creates a new Folder instance
func NewFolder(path string) Folder {
	return Folder(path)
}

// Name returns the name of the folder
func (f Folder) Name() string {
	return filepath.Base(string(f))
}

// Path returns the path of the folder
func (f Folder) Path() string {
	return string(f)
}

// Create creates the folder if it doesn't exist
func (f Folder) Create() error {
	return os.MkdirAll(string(f), os.ModePerm)
}

// Exists checks if the folder exists
func (f Folder) Exists() bool {
	info, err := os.Stat(string(f))
	return err == nil && info.IsDir()
}

// Ls lists the contents of the folder
func (f Folder) Ls() ([]fs.DirEntry, error) {
	return os.ReadDir(string(f))
}

// WriteFile writes data to a file in the folder
func (f Folder) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filepath.Join(string(f), name), data, perm)
}

// ReadFile reads the contents of a file in the folder
func (f Folder) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(string(f), name))
}

// CopyFile copies a file from src to dst within the folder
func (f Folder) CopyFile(src, dst string) error {
	srcPath := filepath.Join(string(f), src)
	dstPath := filepath.Join(string(f), dst)

	sourceFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// DeleteFile deletes a file from the folder
func (f Folder) DeleteFile(name string) error {
	return os.Remove(filepath.Join(string(f), name))
}

// Remove removes the folder and its contents
func (f Folder) Remove() error {
	return os.RemoveAll(string(f))
}

// Join joins the folder path with the given elements
func (f Folder) Join(elem ...string) Folder {
	return Folder(filepath.Join(append([]string{string(f)}, elem...)...))
}

// IsDir checks if the folder is a directory
func (f Folder) IsDir() bool {
	info, err := os.Stat(string(f))
	return err == nil && info.IsDir()
}
