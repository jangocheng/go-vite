package common

import (
	"path/filepath"
	"os"
	"os/user"
	"runtime"
)

// DefaultDataDir is  $HOME/viteisbest/
func DefaultDataDir() string {
	home := HomeDir()
	if home != "" {
		return filepath.Join(home, "viteisbest")
	}
	return ""
}

//it is the dir in go-vite/testdata
func GoViteTestDataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filepath.Dir(filename)), "testdata")
}

func HomeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}

