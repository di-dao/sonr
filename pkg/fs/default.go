package fs

import (
	"os"
	"path/filepath"
)

const kVaultsFolderName = ".sonr-vaults"

var VaultsFolder Folder

func init() {
	// Initialize VaultsFolder
	f, err := NewFolder(kVaultsFolderName)
	if err != nil {
		panic(err)
	}
	VaultsFolder = f
}

func FetchVaultPath(name string) string {
	return filepath.Join(os.ExpandEnv("$HOME"), kVaultsFolderName, name)
}
