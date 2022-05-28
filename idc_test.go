// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppIDCEnv(t *testing.T) {
	t.Run("init with value", func(t *testing.T) {
		e1 := NewAppIDCEnv("demo")
		require.Equal(t, "demo", e1.IDC())
		t.Run("twice", func(t *testing.T) {
			defer func() {
				require.NotNil(t, recover())
			}()
			e1.SetIDC("test")
		})
	})

	t.Run("init with empty", func(t *testing.T) {
		e1 := NewAppIDCEnv("")
		require.NoError(t, os.Setenv(eKeyIDC, "test"))
		defer func() {
			require.NoError(t, os.Unsetenv(eKeyLog))
		}()
		require.Equal(t, "test", e1.IDC())
		e1.SetIDC("v2")
		require.Equal(t, "v2", e1.IDC())
	})
}
