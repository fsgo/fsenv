/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/5/2
 */

package fsenv

// IAppLogEnv 日志目录环境信息
type IAppLogEnv interface {
	LogRootDir() string
	SetLogRootDir(dir string)
}

// NewAppLogEnv 创建新的日志目录环境信息
func NewAppLogEnv(dir string) IAppLogEnv {
	return &LogEnv{
		logDir: dir,
	}
}

type LogEnv struct {
	logDir  string
	rootEnv IAppRootEnv
}

func (l *LogEnv) LogRootDir() string {
	return chooseDirWithRootEnv(l.logDir, l.rootEnv, "log")
}

func (l *LogEnv) SetLogRootDir(dir string) {
	setOnce(&l.logDir, dir, "LogDir")
}

var _ IAppLogEnv = (*LogEnv)(nil)
