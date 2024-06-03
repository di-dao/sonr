package ipfs

import (
	"fmt"
	"io"

	"github.com/di-dao/sonr/crypto/kss"
	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
)

// VFS is an interface for interacting with a virtual file system.
type VFS interface {
	Add(path string, data []byte) error
	AddFileMap(files map[string]files.File) error
	Get(path string) ([]byte, error)
	Rm(path string) error
	Ls(path string) ([]string, error)
	Name() string
	Node() files.Node
	Path() path.Path
	Validate() error

	setPath(path path.Path)
}

// vfs is the struct implementation of an IPFS file system
type vfs struct {
	path  path.Path
	files map[string]files.File
	name  string
}

// NewVFS creates a new virtual file system.
func NewFS(name string) VFS {
	return &vfs{
		files: make(map[string]files.File, 0),
		name:  name,
	}
}

// NewFSWithKss creates a new virtual file system from a kss setPath
func NewFSWithKss(kss kss.Set, name string) VFS {
	fs := &vfs{
		files: make(map[string]files.File, 0),
		name:  name,
	}

	return fs
}

// Load creates a new virtual file system from a given files.Node.
func Load(name string, node files.Node) (VFS, error) {
	entry := files.FileEntry(name, node)
	rootDir := files.DirFromEntry(entry)
	vfs := &vfs{
		files: make(map[string]files.File, 0),
		name:  name,
	}

	err := loadDirectory(rootDir, vfs)
	if err != nil {
		return nil, err
	}

	return vfs, nil
}

// loadDirectory recursively loads the files and directories from a given directory node.
func loadDirectory(dir files.Directory, vfs *vfs) error {
	it := dir.Entries()
	for it.Next() {
		name, node := it.Name(), it.Node()
		switch node := node.(type) {
		case files.File:
			data, err := io.ReadAll(node)
			if err != nil {
				return err
			}
			vfs.files[name] = files.NewBytesFile(data)

		case files.Directory:
			subDir := files.DirFromEntry(files.FileEntry(name, node))
			err := loadDirectory(subDir, vfs)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported node type")
		}
	}
	return nil
}

// Name returns the name of the virtual file system.
func (v *vfs) Name() string {
	return v.name
}

// Add adds a file to the virtual file system.
func (v *vfs) Add(path string, data []byte) error {
	v.files[path] = files.NewBytesFile(data)
	return nil
}

// AddFileMap adds a file map to the virtual file system
func (v *vfs) AddFileMap(files map[string]files.File) error {
	for k, f := range files {
		v.files[k] = f
	}
	return nil
}

// Get retrieves a file from the virtual file system.
func (v *vfs) Get(path string) ([]byte, error) {
	if file, ok := v.files[path]; ok {
		return io.ReadAll(file)
	}
	return nil, fmt.Errorf("file not found")
}

// Rm removes a file from the virtual file system.
func (v *vfs) Rm(path string) error {
	delete(v.files, path)
	return nil
}

// Ls lists the files in the virtual file system.
func (v *vfs) Ls(path string) ([]string, error) {
	var files []string
	for k := range v.files {
		files = append(files, k)
	}
	return files, nil
}

// Node returns the root node of the virtual file system.
func (v *vfs) Node() files.Node {
	rootDir := make(map[string]files.Node, 0)
	fileList := make([]files.DirEntry, 0)
	for k, f := range v.files {
		ent := files.FileEntry(k, f)
		fileList = append(fileList, ent)
	}
	dir := files.NewSliceDirectory(fileList)
	node := dir.Entries().Node()
	rootDir[v.name] = node
	finalDir := files.NewMapDirectory(rootDir)
	return finalDir.Entries().Node()
}

// Path returns the path for the virtual file system
func (v *vfs) Path() path.Path {
	return v.path
}

// Validate validates the virtual file system and returns an error if it is invalid
func (v *vfs) Validate() error {
	if v.name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if v.path == nil {
		return fmt.Errorf("path cannot be empty")
	}
	return nil
}

// setPath sets the path for the virtual file system
func (v *vfs) setPath(path path.Path) {
	v.path = path
}
