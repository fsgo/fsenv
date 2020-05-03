/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/5/2
 */

package fsenv

import (
	"fmt"
	"log"
	"path/filepath"
)

// Value 具体的环境信息
type Value struct {
	RootDir string
	DataDir string
	LogDir  string
	ConfDir string
}

// AppEnv 应用环境
type AppEnv interface {
	AppRootEnv
	AppDataEnv
	AppLogEnv
	AppConfEnv

	Value() *Value
}

// NewAppEnv 创建新的应用环境
func NewAppEnv(opt *Value) AppEnv {
	if opt == nil {
		opt = &Value{}
	}
	root := opt.RootDir
	if root == "" {
		root = AppRootDir()
	}
	env := &appEnv{}
	env.SetRootDir(root)
	env.SetDataRootDir(choose(opt.DataDir, filepath.Join(root, "data")))
	env.SetLogRootDir(choose(opt.LogDir, filepath.Join(root, "log")))
	env.SetConfRootPath(choose(opt.ConfDir, filepath.Join(root, "conf")))
	return env
}

func choose(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

type appEnv struct {
	*rootEnv
	*dataEnv
	*logEnv
	*confEnv
}

func (e *appEnv) Value() *Value {
	return &Value{
		RootDir: e.RootDir(),
		DataDir: e.DataRootDir(),
		LogDir:  e.LogRootDir(),
		ConfDir: e.ConfRootPath(),
	}
}

var _ AppEnv = (*appEnv)(nil)

func setOnce(addr *string, value string, fieldName string) {
	if *addr != "" {
		panic(fmt.Sprintf("cannot set %s twice", fieldName))
	}
	*addr = value
	log.Output(2, fmt.Sprintf("[fsenv] set %q=%q\n", fieldName, value))
}
