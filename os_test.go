// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv

import (
	"os"
	"testing"

	"github.com/fsgo/fst"
)

func Test_osEnvDefault(t *testing.T) {
	key := "fsenv_k1"
	fst.NoError(t, os.Unsetenv(key))
	defer fst.NoError(t, os.Unsetenv(key))

	fst.Equal(t, "v1", osEnvDefault(key, "v1"))
	fst.NoError(t, os.Setenv(key, "v2"))
	fst.Equal(t, "v2", osEnvDefault(key, "v1"))
}
