// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/2

package appenv

import (
	"fmt"
	"os"
	"path/filepath"
)

// AppRoot 自动获取应用根目录
func AppRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	names := []string{
		"go.mod",
	}
	dir, err := findDirMatch(wd, names)
	if err == nil {
		return dir
	}
	return wd
}

var errNotFound = fmt.Errorf("cannot found")

func findDirMatch(baseDir string, fileNames []string) (dir string, err error) {
	currentDir := baseDir
	for i := 0; i < 50; i++ {
		for _, fileName := range fileNames {
			depsPath := filepath.Join(currentDir, fileName)
			if _, err := os.Stat(depsPath); !os.IsNotExist(err) {
				return currentDir, nil
			}
		}

		currentDir = filepath.Dir(currentDir)

		if currentDir == "." {
			break
		}
	}
	return "", errNotFound
}
