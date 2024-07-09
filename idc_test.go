// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv

import (
	"os"
	"testing"

	"github.com/fsgo/fst"
)

func TestAppIDCEnv(t *testing.T) {
	t.Run("init with value", func(t *testing.T) {
		e1 := NewAppIDCEnv("demo")
		fst.Equal(t, "demo", e1.IDC())
		t.Run("twice", func(t *testing.T) {
			defer func() {
				fst.Nil(t, recover())
			}()
			e1.SetIDC("test")
		})
	})

	t.Run("init with empty", func(t *testing.T) {
		e1 := NewAppIDCEnv("")
		fst.NoError(t, os.Setenv(eKeyIDC, "test"))
		defer func() {
			fst.NoError(t, os.Unsetenv(eKeyLog))
		}()
		fst.Equal(t, "test", e1.IDC())
		e1.SetIDC("v2")
		fst.Equal(t, "v2", e1.IDC())
	})
}
