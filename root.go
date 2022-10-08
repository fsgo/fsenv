// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/2

package fsenv

import (
	"os"
	"path/filepath"

	"github.com/fsgo/fsenv/internal/appenv"
)

// AppRootDir 自动推断、获取应用根目录
//
//	 推断顺序：
//		1. 环境变量：fsenv.root
//		2. 查找 go.mod 文件
//		3. 当前目录（pwd）
func AppRootDir() string {
	def := os.Getenv(eKeyRoot)
	if len(def) == 0 {
		return appenv.AppRoot()
	}
	return def
}

// AppRootEnv 应用更目录环境信息
type AppRootEnv interface {
	HasRootDir
	CanSetRootDir
}

// HasRootDir 可以获取根目录
type HasRootDir interface {
	RootDir() string
}

// CanSetRootDir 可以设置根目录
type CanSetRootDir interface {
	SetRootDir(dir string)
}

// NewAppRootEnv 创建新的应用更目录环境
func NewAppRootEnv(root string) AppRootEnv {
	return &rootEnv{
		rootDir: root,
	}
}

type rootEnv struct {
	rootDir string
}

func (r *rootEnv) RootDir() string {
	if len(r.rootDir) > 0 {
		return r.rootDir
	}
	return AppRootDir()
}

func (r *rootEnv) SetRootDir(dir string) {
	setOnce(&r.rootDir, dir, "RootDir")
}

var _ AppRootEnv = (*rootEnv)(nil)

func chooseDirWithRootEnv(dir string, env HasRootDir, key string, subDirName string) string {
	if len(dir) > 0 {
		return dir
	}
	// 当没有明确设置的时候，第 2 优先级为环境变量的值
	if val := os.Getenv(key); len(val) > 0 {
		return val
	}
	if env == nil {
		return filepath.Join(AppRootDir(), subDirName)
	}
	return filepath.Join(env.RootDir(), subDirName)
}
