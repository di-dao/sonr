package fs_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ipfs/boxo/files"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/yourusername/yourproject/pkg/fs"
)

func TestFolder(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "folder_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	folder := fs.Folder(tempDir)

	t.Run("Path", func(t *testing.T) {
		assert.Equal(t, tempDir, folder.Path())
	})

	t.Run("Create", func(t *testing.T) {
		newFolder := folder.Join("new_folder")
		err := newFolder.Create()
		assert.NoError(t, err)
		assert.True(t, newFolder.Exists())
	})

	t.Run("Exists", func(t *testing.T) {
		assert.True(t, folder.Exists())
		assert.False(t, folder.Join("non_existent").Exists())
	})

	t.Run("Remove", func(t *testing.T) {
		newFolder := folder.Join("to_remove")
		err := newFolder.Create()
		require.NoError(t, err)

		err = newFolder.Remove()
		assert.NoError(t, err)
		assert.False(t, newFolder.Exists())
	})

	t.Run("ReadDir", func(t *testing.T) {
		err := os.WriteFile(filepath.Join(tempDir, "test.txt"), []byte("test"), 0644)
		require.NoError(t, err)

		entries, err := folder.ReadDir()
		assert.NoError(t, err)
		assert.Len(t, entries, 1)
		assert.Equal(t, "test.txt", entries[0].Name())
	})

	t.Run("Join", func(t *testing.T) {
		joined := folder.Join("subdir", "file.txt")
		expected := filepath.Join(tempDir, "subdir", "file.txt")
		assert.Equal(t, expected, joined.Path())
	})

	t.Run("IsDir", func(t *testing.T) {
		assert.True(t, folder.IsDir())
		assert.False(t, folder.Join("test.txt").IsDir())
	})

	t.Run("Node", func(t *testing.T) {
		node, err := folder.Node()
		assert.NoError(t, err)
		assert.IsType(t, &files.MapDirectory{}, node)
	})

	t.Run("LoadNodeInFolder", func(t *testing.T) {
		content := []byte("test content")
		fileNode := files.NewBytesFile(content)
		dirMap := map[string]files.Node{
			"test.txt": fileNode,
		}
		node := files.NewMapDirectory(dirMap)

		loadPath := filepath.Join(tempDir, "loaded")
		loadedFolder, err := fs.LoadNodeInFolder(loadPath, node)
		assert.NoError(t, err)
		assert.True(t, loadedFolder.Exists())

		loadedContent, err := os.ReadFile(filepath.Join(loadPath, "test.txt"))
		assert.NoError(t, err)
		assert.Equal(t, content, loadedContent)
	})
}
