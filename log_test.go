// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv

import (
	"os"
	"testing"

	"github.com/fsgo/fst"
)

func TestAppLogEnv(t *testing.T) {
	t.Run("init with value", func(t *testing.T) {
		e1 := NewAppLogEnv("demo")
		fst.Equal(t, "demo", e1.LogRootDir())
		t.Run("twice", func(t *testing.T) {
			defer func() {
				fst.Nil(t, recover())
			}()
			e1.SetLogRootDir("test")
		})
	})

	t.Run("init with empty", func(t *testing.T) {
		e1 := NewAppLogEnv("")
		fst.NoError(t, os.Setenv(eKeyLog, "test"))
		defer func() {
			fst.NoError(t, os.Unsetenv(eKeyLog))
		}()
		fst.Equal(t, "test", e1.LogRootDir())
		e1.SetLogRootDir("v2")
		fst.Equal(t, "v2", e1.LogRootDir())
	})
}
