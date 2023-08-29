package local

import (
	"os"
	"path/filepath"
)

// LocalContext is a struct that holds the current context of the application.
type LocalContext struct {
	HomeDir        string
	NodeHome       string
	ConfigDirPath  string
	ConfigTomlPath string
}

// Option is a function that configures the local context
type Option func(LocalContext)

// Context returns the current context of the Sonr blockchain application.
func Context(opts ...Option) LocalContext {
	c := LocalContext{
		HomeDir:        filepath.Join(HomeDir()),
		NodeHome:       filepath.Join(HomeDir(), ".sonr"),
		ConfigDirPath:  filepath.Join(HomeDir(), ".sonr", "config"),
		ConfigTomlPath: filepath.Join(HomeDir(), ".sonr", "config", "config.toml"),
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

// The HomeDir function returns the home directory of the current user.
func HomeDir() string {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		homeDir = os.Getenv("USERPROFILE") // windows
	}
	return homeDir
}
