package fs

import (
	"os"
	"path/filepath"
)

const kVaultsFolderName = ".sonr-vaults"

var VaultsFolder Folder

func init() {
	// Initialize VaultsFolder
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	VaultsFolder = NewFolder(filepath.Join(homeDir, kVaultsFolderName))
	if !VaultsFolder.Exists() {
		if err := VaultsFolder.Create(); err != nil {
			panic(err)
		}
	}
}

func FetchVaultPath(name string) string {
	return VaultsFolder.Join(name).Path()
}
