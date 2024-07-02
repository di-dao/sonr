package fs

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/ipfs/boxo/path"
)

// Helper function to parse IPFS path
func ParsePath(p string) (path.Path, error) {
	return path.NewPath(p)
}

// FetchVaultPath returns the path to the vault with the given name
func FetchVaultPath(name string) string {
	return VaultsFolder.Join(name).Path()
}

// OpenURL is a Helper function which opens a url in windows,linux,osx default browser
func OpenURL(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
