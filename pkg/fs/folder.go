package fs

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/ipfs/boxo/files"
)

type Folder string

func (f Folder) Path() string {
	return string(f)
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

// Node returns the folder as an IPFS node
func (f Folder) Node() (files.Node, error) {
	return f.loadDirectory(f.Path())
}

// loadDirectory recursively loads the files and directories from a given path
func (f Folder) loadDirectory(path string) (files.Node, error) {
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

// Load creates a new Folder from a given files.Node
func LoadNodeInFolder(path string, node files.Node) (Folder, error) {
	folder := Folder(path)
	err := folder.Create()
	if err != nil {
		return "", err
	}

	err = folder.loadFromNode(node, path)
	if err != nil {
		return "", err
	}

	return folder, nil
}

// loadFromNode recursively loads files and directories from a given files.Node
func (f Folder) loadFromNode(node files.Node, path string) error {
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
