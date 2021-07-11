// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/2

package fsenv

// AppDataEnv 数据目录环境信息
type AppDataEnv interface {
	DataRootDir() string
	SetDataRootDir(dir string)
}

// NewAppDataEnv 创建新的数据目录环境
func NewAppDataEnv(dir string) AppDataEnv {
	return &dataEnv{
		dataDir: dir,
	}
}

type dataEnv struct {
	dataDir string
	rootEnv AppRootEnv
}

func (d *dataEnv) DataRootDir() string {
	return chooseDirWithRootEnv(d.dataDir, d.rootEnv, "data")
}

func (d *dataEnv) SetDataRootDir(dir string) {
	setOnce(&d.dataDir, dir, "DataDir")
}

var _ AppDataEnv = (*dataEnv)(nil)
