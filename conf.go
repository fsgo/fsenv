// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/2

package fsenv

// AppConfEnv 配置环境信息
type AppConfEnv interface {
	HasConfRootDir
	CanSetConfRootDir
}

// HasConfRootDir 可以获取配置目录
type HasConfRootDir interface {
	ConfRootDir() string
}

// CanSetConfRootDir 运行设置配置目录
type CanSetConfRootDir interface {
	SetConfRootDir(dir string)
}

// NewAppConfEnv 创建新的配置环境信息
func NewAppConfEnv(dir string) AppConfEnv {
	return &confEnv{
		confDir: dir,
	}
}

type confEnv struct {
	rootEnv AppRootEnv
	confDir string
}

func (c *confEnv) ConfRootDir() string {
	return chooseDirWithRootEnv(c.confDir, c.rootEnv, eKeyConf, "conf")
}

func (c *confEnv) SetConfRootDir(dir string) {
	setOnce(&c.confDir, dir, "ConfDir")
}

var _ AppConfEnv = (*confEnv)(nil)
