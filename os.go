// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv

import (
	"os"
)

const (
	eKeyIDC  = "fsenv_idc"
	eKeyConf = "fsenv_conf"
	eKeyRoot = "fsenv_root"
	eKeyData = "fsenv_data"
	eKeyLog  = "fsenv_log"
	eKeyMode = "fsenv_mode"
)

func osEnvDefault(key string, def string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return def
	}
	return val
}
