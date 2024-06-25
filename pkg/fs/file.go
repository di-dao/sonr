package fs

import (
	"io"
	"os"
	"time"
)

// File represents a file in the filesystem
type File struct {
	name string
	data []byte
}

// NewFile creates a new File instance
func NewFile(name string, data []byte) *File {
	return &File{
		name: name,
		data: data,
	}
}

// Name returns the name of the file
func (f *File) Name() string {
	return f.name
}

// Read implements io.Reader
func (f *File) Read(p []byte) (n int, err error) {
	return copy(p, f.data), io.EOF
}

// Close implements io.Closer
func (f *File) Close() error {
	return nil
}

// Stat implements os.FileInfo
func (f *File) Stat() (os.FileInfo, error) {
	return &fileInfo{f}, nil
}

// fileInfo implements os.FileInfo
type fileInfo struct {
	file *File
}

func (fi *fileInfo) Name() string       { return fi.file.name }
func (fi *fileInfo) Size() int64        { return int64(len(fi.file.data)) }
func (fi *fileInfo) Mode() os.FileMode  { return 0644 }
func (fi *fileInfo) ModTime() time.Time { return time.Now() }
func (fi *fileInfo) IsDir() bool        { return false }
func (fi *fileInfo) Sys() interface{}   { return nil }
