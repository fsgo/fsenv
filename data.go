// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/2

package fsenv

// AppDataEnv 数据目录环境信息
type AppDataEnv interface {
	HasDataRootDir
	CanSetDataRootDir
}

// HasDataRootDir 可以获取数据目录
type HasDataRootDir interface {
	DataRootDir() string
}

// CanSetDataRootDir 允许设置数据目录
type CanSetDataRootDir interface {
	SetDataRootDir(dir string)
}

// NewAppDataEnv 创建新的数据目录环境
func NewAppDataEnv(dir string) AppDataEnv {
	return &dataEnv{
		dataDir: dir,
	}
}

type dataEnv struct {
	rootEnv AppRootEnv
	dataDir string
}

func (d *dataEnv) DataRootDir() string {
	return chooseDirWithRootEnv(d.dataDir, d.rootEnv, eKeyData, "data")
}

func (d *dataEnv) SetDataRootDir(dir string) {
	setValue(&d.dataDir, dir, "DataDir")
}

var _ AppDataEnv = (*dataEnv)(nil)
