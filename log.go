// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/2

package fsenv

// AppLogEnv 日志目录环境信息
type AppLogEnv interface {
	LogRootDir() string
	SetLogRootDir(dir string)
}

// NewAppLogEnv 创建新的日志目录环境信息
func NewAppLogEnv(dir string) AppLogEnv {
	return &logEnv{
		logDir: dir,
	}
}

type logEnv struct {
	logDir  string
	rootEnv AppRootEnv
}

func (l *logEnv) LogRootDir() string {
	return chooseDirWithRootEnv(l.logDir, l.rootEnv, "log")
}

func (l *logEnv) SetLogRootDir(dir string) {
	setOnce(&l.logDir, dir, "LogDir")
}

var _ AppLogEnv = (*logEnv)(nil)
