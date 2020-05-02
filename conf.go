/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/5/2
 */

package fsenv

// AppConfEnv 配置环境信息
type AppConfEnv interface {
	ConfRootPath() string
	SetConfRootPath(dir string)
}

// NewAppConfEnv 创建新的配置环境信息
func NewAppConfEnv(dir string) AppConfEnv {
	return &confEnv{
		confDir: dir,
	}
}

type confEnv struct {
	confDir string
	rootEnv AppRootEnv
}

func (c *confEnv) ConfRootPath() string {
	return chooseDirWithRootEnv(c.confDir, c.rootEnv, "conf")
}

func (c *confEnv) SetConfRootPath(dir string) {
	setOnce(&c.confDir, dir, "ConfDir")
}

var _ AppConfEnv = (*confEnv)(nil)
