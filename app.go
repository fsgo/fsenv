// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/2

package fsenv

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

// Value 具体的环境信息
type Value struct {
	RootDir string
	DataDir string
	LogDir  string
	ConfDir string
	IDC     string
	RunMode Mode
}

// String 序列化，方便查看
func (v *Value) String() string {
	bf, _ := json.Marshal(v)
	return string(bf)
}

// AppEnv 应用环境信息完整的接口定义
type AppEnv interface {
	RootDir() string
	DataRootDir() string
	LogRootDir() string
	ConfRootDir() string
	IDC() string
	RunMode() Mode
	Value() Value
}

// AppEnvRW 可读写的接口定义
type AppEnvRW interface {
	AppRootEnv
	AppDataEnv
	AppLogEnv
	AppConfEnv
	AppIDCEnv
	AppRunMode

	Value() Value
}

// NewAppEnv 创建新的应用环境
func NewAppEnv(opt Value) AppEnv {
	root := opt.RootDir
	if len(root) == 0 {
		root = AppRootDir()
	}
	envRoot := &rootEnv{
		rootDir: root,
	}

	mode := ModeProduct
	if len(opt.RunMode) != 0 {
		mode = opt.RunMode
	}
	env := &appEnv{
		AppRootEnv: envRoot,
		AppDataEnv: &dataEnv{
			rootEnv: envRoot,
			dataDir: choose(opt.DataDir, filepath.Join(root, "data")),
		},
		AppLogEnv: &logEnv{
			rootEnv: envRoot,
			logDir:  choose(opt.LogDir, filepath.Join(root, "log")),
		},
		AppConfEnv: &confEnv{
			rootEnv: envRoot,
			confDir: choose(opt.ConfDir, filepath.Join(root, "conf")),
		},
		AppIDCEnv: &idcEnv{
			idc: choose(opt.IDC, "test"),
		},
		AppRunMode: NewAppRunModeEnv(mode),
	}
	return env
}

func choose(a, b string) string {
	if len(a) != 0 {
		return a
	}
	return b
}

// appEnv 默认实现的应用环境信息
type appEnv struct {
	AppRootEnv
	AppDataEnv
	AppLogEnv
	AppConfEnv
	AppIDCEnv
	AppRunMode
}

// Value 获取环境信息所有的值
func (e *appEnv) Value() Value {
	return Value{
		RootDir: e.RootDir(),
		DataDir: e.DataRootDir(),
		LogDir:  e.LogRootDir(),
		ConfDir: e.ConfRootDir(),
		IDC:     e.IDC(),
		RunMode: e.RunMode(),
	}
}

var _ AppEnv = (*appEnv)(nil)

func setOnce(addr *string, value string, fieldName string) {
	if len(*addr) > 0 {
		panic(fmt.Sprintf("cannot set %s twice", fieldName))
	}
	*addr = value
}
