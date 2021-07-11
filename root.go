// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/2

package fsenv

import (
	"path/filepath"

	"github.com/fsgo/fsenv/internal/appenv"
)

// AppRootDir 自动推断、获取应用根目录
// 目前是采用查找go.mod文件
// 若查找不到则返回当前目录（pwd）
func AppRootDir() string {
	return appenv.AppRoot()
}

// AppRootEnv 应用更目录环境信息
type AppRootEnv interface {
	RootDir() string
	SetRootDir(dir string)
}

// NewAppRootEnv 创建新的应用更目录环境
func NewAppRootEnv(root string) AppRootEnv {
	if root == "" || root == "auto" {
		root = AppRootDir()
	}
	return &rootEnv{
		rootDir: root,
	}
}

type rootEnv struct {
	rootDir string
}

func (r *rootEnv) RootDir() string {
	if r.rootDir != "" {
		return r.rootDir
	}
	return AppRootDir()
}

func (r *rootEnv) SetRootDir(dir string) {
	setOnce(&r.rootDir, dir, "RootDir")
}

var _ AppRootEnv = (*rootEnv)(nil)

func chooseDirWithRootEnv(dir string, env AppRootEnv, subDirName string) string {
	if dir != "" {
		return dir
	}
	if env == nil {
		return filepath.Join(AppRootDir(), subDirName)
	}
	return filepath.Join(env.RootDir(), subDirName)
}
