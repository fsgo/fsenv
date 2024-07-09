// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv

import (
	"os"
	"testing"

	"github.com/fsgo/fst"
)

func TestAppConfEnv(t *testing.T) {
	t.Run("init with value", func(t *testing.T) {
		e1 := NewAppConfEnv("demo")
		fst.Equal(t, "demo", e1.ConfRootDir())
		t.Run("twice", func(t *testing.T) {
			defer func() {
				fst.Nil(t, recover())
			}()
			e1.SetConfRootDir("test")
		})
	})

	t.Run("init with empty", func(t *testing.T) {
		e1 := NewAppConfEnv("")
		fst.NoError(t, os.Setenv(eKeyConf, "test"))
		defer func() {
			fst.NoError(t, os.Unsetenv(eKeyConf))
		}()
		fst.Equal(t, "test", e1.ConfRootDir())
		e1.SetConfRootDir("v2")
		fst.Equal(t, "v2", e1.ConfRootDir())
	})
}
