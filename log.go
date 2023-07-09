// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/2

package fsenv

// AppLogEnv 日志目录环境信息
type AppLogEnv interface {
	HasLogRootDir
	CanSetLogRootDir
}

// HasLogRootDir 可以获取日志根目录
type HasLogRootDir interface {
	LogRootDir() string
}

// CanSetLogRootDir 可以设置日志根目录
type CanSetLogRootDir interface {
	SetLogRootDir(dir string)
}

// NewAppLogEnv 创建新的日志目录环境信息
func NewAppLogEnv(dir string) AppLogEnv {
	return &logEnv{
		logDir: dir,
	}
}

type logEnv struct {
	rootEnv AppRootEnv
	logDir  string
}

func (l *logEnv) LogRootDir() string {
	return chooseDirWithRootEnv(l.logDir, l.rootEnv, eKeyLog, "log")
}

func (l *logEnv) SetLogRootDir(dir string) {
	setValue(&l.logDir, dir, "LogDir")
}

var _ AppLogEnv = (*logEnv)(nil)
