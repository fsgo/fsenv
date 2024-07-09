// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv

import (
	"os"
	"testing"

	"github.com/fsgo/fst"
)

func TestAppDataEnv(t *testing.T) {
	t.Run("init with value", func(t *testing.T) {
		e1 := NewAppDataEnv("demo")
		fst.Equal(t, "demo", e1.DataRootDir())
		t.Run("twice", func(t *testing.T) {
			defer func() {
				fst.Nil(t, recover())
			}()
			e1.SetDataRootDir("test")
		})
	})

	t.Run("init with empty", func(t *testing.T) {
		e1 := NewAppDataEnv("")
		fst.NoError(t, os.Setenv(eKeyData, "test"))
		defer func() {
			fst.NoError(t, os.Unsetenv(eKeyData))
		}()
		fst.Equal(t, "test", e1.DataRootDir())
		e1.SetDataRootDir("v2")
		fst.Equal(t, "v2", e1.DataRootDir())
	})
}
