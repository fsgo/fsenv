// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv

import (
	"os"
)

const (
	eKeyIDC  = "fsenv.idc"
	eKeyConf = "fsenv.conf"
	eKeyRoot = "fsenv.root"
	eKeyData = "fsenv.data"
	eKeyLog  = "fsenv.log"
	eKeyMode = "fsenv.mode"
)

func osEnvDefault(key string, def string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return def
	}
	return val
}
