package device

import (
	"errors"
	"os"
	"runtime"

	"github.com/denisbrodbeck/machineid"
)

// DeviceStat is the device info struct
type DeviceStat struct {
	Id        string `json:"id"`
	HostName  string `json:"name"`
	Os        string `json:"os"`
	Arch      string `json:"arch"`
	IsDesktop bool
	IsMobile  bool
}

// AppName returns the application name.
func AppName() string {
	switch runtime.GOOS {
	case "android":
		return "io.sonr.petitfour"
	case "darwin":
		return "io.sonr.macos"
	case "linux":
		return "io.sonr.linux"
	case "windows":
		return "io.sonr.windows"
	case "ios":
		return "io.sonr.alpine"
	default:
		return "io.sonr.app"
	}
}

// Arch returns the current architecture.
func Arch() string {
	return runtime.GOARCH
}

// HostName returns the hostname of the current machine.
func HostName() (string, error) {
	return os.Hostname()
}

// ID returns the device ID.
func ID() (string, error) {
	// Check if Mobile
	if IsMobile() {
		if deviceID != "" {
			return deviceID, nil
		}
		return "", errors.New("Device ID not set for Mobile.")
	}
	return machineid.ID()
}

// IsMobile returns true if the current platform is ANY mobile platform.
func IsMobile() bool {
	return runtime.GOOS == "android" || runtime.GOOS == "ios"
}

// IsDesktop returns true if the current platform is ANY desktop platform.
func IsDesktop() bool {
	return runtime.GOOS == "windows" || runtime.GOOS == "linux" || runtime.GOOS == "darwin"
}

// IsAndroid returns true if the current platform is android.
func IsAndroid() bool {
	return runtime.GOOS == "android"
}

// IsIOS returns true if the current platform is iOS.
func IsIOS() bool {
	return runtime.GOOS == "ios"
}

// IsWindows returns true if the current platform is windows.
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux returns true if the current platform is linux.
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsMacOS returns true if the current platform is macOS.
func IsMacOS() bool {
	return runtime.GOOS == "darwin"
}

// VendorName returns the vendor name.
func VendorName() string {
	return "Sonr"
}

// Stat returns the device stat.
func Stat() (*DeviceStat, error) {
	// Get Device Id
	id, err := ID()
	if err != nil {
		logger.Error("Failed to get Device ID", err)
		return nil, err
	}

	// Get HostName
	hn, err := HostName()
	if err != nil {
		logger.Error("Failed to get HostName", err)
		return nil, err
	}

	// Return the device info for Peer
	return &DeviceStat{
		Id:        id,
		HostName:  hn,
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		IsDesktop: IsDesktop(),
		IsMobile:  IsMobile(),
	}, nil
}
