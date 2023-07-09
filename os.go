// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv

import (
	"os"
)

const (
	eKeyIDC  = "FSENV_IDC"
	eKeyConf = "FSENV_CONF"
	eKeyRoot = "FSENV_ROOT"
	eKeyData = "FSENV_DATA"
	eKeyLog  = "FSENV_LOG"
	eKeyMode = "FSENV_MODE"
)

func osEnvDefault(key string, def string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return def
	}
	return val
}
