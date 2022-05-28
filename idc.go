// Copyright(C) 2021 github.com/fsgo  All Rights Reserved.
// Author: fsgo
// Date: 2021/7/11

package fsenv

// AppIDCEnv 机房信息
type AppIDCEnv interface {
	HasIDC
	CanSetIDC
}

// HasIDC 可以获取 idc
type HasIDC interface {
	IDC() string
}

// CanSetIDC 运行设置 idc
type CanSetIDC interface {
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
	if len(c.idc) > 0 {
		return c.idc
	}
	return osEnvDefault(eKeyIDC, "test")
}

func (c *idcEnv) SetIDC(idc string) {
	setOnce(&c.idc, idc, "IDC")
}

var _ AppIDCEnv = (*idcEnv)(nil)
