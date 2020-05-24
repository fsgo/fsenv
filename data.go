/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/5/2
 */

package fsenv

// IAppDataEnv 数据目录环境信息
type IAppDataEnv interface {
	DataRootDir() string
	SetDataRootDir(dir string)
}

// NewAppDataEnv 创建新的数据目录环境
func NewAppDataEnv(dir string) IAppDataEnv {
	return &DataEnv{
		dataDir: dir,
	}
}

type DataEnv struct {
	dataDir string
	rootEnv IAppRootEnv
}

func (d *DataEnv) DataRootDir() string {
	return chooseDirWithRootEnv(d.dataDir, d.rootEnv, "data")
}

func (d *DataEnv) SetDataRootDir(dir string) {
	setOnce(&d.dataDir, dir, "DataDir")
}

var _ IAppDataEnv = (*DataEnv)(nil)
