// Copyright(C) 2021 github.com/fsgo  All Rights Reserved.
// Author: fsgo
// Date: 2021/7/11

package fsenv

// AppIDCEnv 机房信息
type AppIDCEnv interface {
	IDC() string
	SetIDC(idc string)
}

// NewAppIDCEnv 创建新的配置环境信息
func NewAppIDCEnv(idc string) AppIDCEnv {
	return &idcEnv{
		idc: idc,
	}
}

type idcEnv struct {
	idc string
}

func (c *idcEnv) IDC() string {
	if c.idc != "" {
		return c.idc
	}
	return "test"
}

func (c *idcEnv) SetIDC(idc string) {
	setOnce(&c.idc, idc, "IDC")
}

var _ AppIDCEnv = (*idcEnv)(nil)
