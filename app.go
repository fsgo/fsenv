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

// String 序列化，方便查看
func (v *Value) String() string {
	format := `{"RootDir":%q,"DataDir":%q,"LogDir":%q,"ConfDir":%q}`
	return fmt.Sprintf(format, v.RootDir, v.DataDir, v.LogDir, v.ConfDir)
}

// IAppEnv 应用环境信息完整的接口定义
type IAppEnv interface {
	IAppRootEnv
	IAppDataEnv
	IAppLogEnv
	IAppConfEnv

	Value() Value
}

// NewAppEnv 创建新的应用环境
func NewAppEnv(opt Value) IAppEnv {
	root := opt.RootDir
	if root == "" {
		root = AppRootDir()
	}
	rootEnv := &RootEnv{
		rootDir: root,
	}

	env := &AppEnv{
		RootEnv: rootEnv,
		DataEnv: &DataEnv{
			rootEnv: rootEnv,
			dataDir: choose(opt.DataDir, filepath.Join(root, "data")),
		},
		LogEnv: &LogEnv{
			rootEnv: rootEnv,
			logDir:  choose(opt.LogDir, filepath.Join(root, "log")),
		},
		ConfEnv: &ConfEnv{
			rootEnv: rootEnv,
			confDir: choose(opt.ConfDir, filepath.Join(root, "conf")),
		},
	}
	return env
}

func choose(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

// AppEnv 默认实现的应用环境信息
type AppEnv struct {
	*RootEnv
	*DataEnv
	*LogEnv
	*ConfEnv
}

// Value 获取环境信息所有的值
func (e *AppEnv) Value() Value {
	return Value{
		RootDir: e.RootDir(),
		DataDir: e.DataRootDir(),
		LogDir:  e.LogRootDir(),
		ConfDir: e.ConfRootDir(),
	}
}

var _ IAppEnv = (*AppEnv)(nil)

func setOnce(addr *string, value string, fieldName string) {
	if *addr != "" {
		panic(fmt.Sprintf("cannot set %s twice", fieldName))
	}
	*addr = value
	log.Output(2, fmt.Sprintf("[fsenv] set %q=%q\n", fieldName, value))
}
