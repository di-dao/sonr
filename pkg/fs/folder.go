package fs

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
)

type Folder struct {
	path path.Path
	name string
}

// NewFolder creates a new Folder instance and creates the folder on disk
func NewFolder(p string) (*Folder, error) {
	folderPath, err := path.NewPath(p)
	if err != nil {
		return nil, fmt.Errorf("failed to create path: %w", err)
	}
	folder := &Folder{
		path: folderPath,
		name: filepath.Base(p),
	}
	err = folder.Create()
	if err != nil {
		return nil, fmt.Errorf("failed to create folder: %w", err)
	}
	return folder, nil
}

// Path returns the path of the folder
func (f *Folder) Path() path.Path {
	return f.path
}

// Name returns the name of the folder
func (f *Folder) Name() string {
	return f.name
}

// SetPath sets the path of the folder
func (f *Folder) SetPath(p string) error {
	newPath, err := path.NewPath(p)
	if err != nil {
		return fmt.Errorf("failed to set path: %w", err)
	}
	f.path = newPath
	f.name = filepath.Base(p)
	return nil
}

// Validate checks if the folder is valid
func (f *Folder) Validate() error {
	if !f.Exists() {
		return fmt.Errorf("folder does not exist: %s", f.path.String())
	}
	return nil
}

// Create creates the folder if it doesn't exist
func (f *Folder) Create() error {
	return os.MkdirAll(f.path.String(), os.ModePerm)
}

// Exists checks if the folder exists
func (f *Folder) Exists() bool {
	_, err := os.Stat(f.path.String())
	return !os.IsNotExist(err)
}

// AddFile adds a file to the directory with the given name and data
func (f *Folder) AddFile(name string, data []byte) error {
	if !f.Exists() {
		if err := f.Create(); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	filePath := filepath.Join(f.path.String(), name)
	return os.WriteFile(filePath, data, 0644)
}

// Remove removes the folder and its contents
func (f *Folder) Remove() error {
	return os.RemoveAll(f.path.String())
}

// ReadDir reads the contents of the folder
func (f *Folder) ReadDir() ([]fs.DirEntry, error) {
	return os.ReadDir(f.path.String())
}

// Join joins the folder path with the given elements
func (f *Folder) Join(elem ...string) (*Folder, error) {
	newPathStr := filepath.Join(append([]string{f.path.String()}, elem...)...)
	newPath, err := path.NewPath(newPathStr)
	if err != nil {
		return nil, fmt.Errorf("failed to create new path: %w", err)
	}
	return &Folder{
		path: newPath,
		name: filepath.Base(newPathStr),
	}, nil
}

// IsDir checks if the folder is a directory
func (f *Folder) IsDir() bool {
	info, err := os.Stat(f.path.String())
	return err == nil && info.IsDir()
}

// Node returns the folder as an IPFS node
func (f *Folder) Node() (files.Node, error) {
	return f.loadDirectory(f.path.String())
}

// loadDirectory recursively loads the files and directories from a given path
func (f *Folder) loadDirectory(path string) (files.Node, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	dirMap := make(map[string]files.Node)

	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())

		if entry.IsDir() {
			subDir, err := f.loadDirectory(fullPath)
			if err != nil {
				return nil, err
			}
			dirMap[entry.Name()] = subDir
		} else {
			file, err := os.Open(fullPath)
			if err != nil {
				return nil, err
			}
			defer file.Close()

			data, err := io.ReadAll(file)
			if err != nil {
				return nil, err
			}
			dirMap[entry.Name()] = files.NewBytesFile(data)
		}
	}

	return files.NewMapDirectory(dirMap), nil
}

// LoadNodeInFolder creates a new Folder from a given files.Node
func LoadNodeInFolder(p string, node files.Node) (*Folder, error) {
	folderPath, err := path.NewPath(p)
	if err != nil {
		return nil, fmt.Errorf("failed to create path: %w", err)
	}
	folder := &Folder{
		path: folderPath,
		name: filepath.Base(p),
	}
	err = folder.Create()
	if err != nil {
		return nil, err
	}

	err = folder.loadFromNode(node, p)
	if err != nil {
		return nil, err
	}

	return folder, nil
}

// loadFromNode recursively loads files and directories from a given files.Node
func (f *Folder) loadFromNode(node files.Node, path string) error {
	switch n := node.(type) {
	case files.File:
		data, err := io.ReadAll(n)
		if err != nil {
			return err
		}
		return os.WriteFile(path, data, 0644)

	case files.Directory:
		it := n.Entries()
		for it.Next() {
			childName := it.Name()
			childNode := it.Node()
			childPath := filepath.Join(path, childName)

			if dir, ok := childNode.(files.Directory); ok {
				err := os.Mkdir(childPath, 0755)
				if err != nil {
					return err
				}
				err = f.loadFromNode(dir, childPath)
				if err != nil {
					return err
				}
			} else if file, ok := childNode.(files.File); ok {
				err := f.loadFromNode(file, childPath)
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("unsupported node type for %s", childName)
			}
		}
		return nil

	default:
		return fmt.Errorf("unsupported node type")
	}
}
