/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/5/2
 */

package fsenv

// IAppConfEnv 配置环境信息
type IAppConfEnv interface {
	ConfRootDir() string
	SetConfRootDir(dir string)
}

// NewAppConfEnv 创建新的配置环境信息
func NewAppConfEnv(dir string) IAppConfEnv {
	return &ConfEnv{
		confDir: dir,
	}
}

type ConfEnv struct {
	confDir string
	rootEnv IAppRootEnv
}

func (c *ConfEnv) ConfRootDir() string {

	return chooseDirWithRootEnv(c.confDir, c.rootEnv, "conf")
}

func (c *ConfEnv) SetConfRootDir(dir string) {
	setOnce(&c.confDir, dir, "ConfDir")
}

var _ IAppConfEnv = (*ConfEnv)(nil)
