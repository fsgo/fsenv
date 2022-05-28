// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/2

package fsenv

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppRoot(t *testing.T) {
	wd, _ := os.Getwd()
	tests := []struct {
		name string
		want string
	}{
		{
			name: "case 1",
			want: wd,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AppRootDir()
			if got != tt.want {
				t.Errorf("AppRootDir() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chooseDirWithRootEnv(t *testing.T) {
	key := "fsenv.k1"
	require.NoError(t, os.Unsetenv(key))
	defer require.NoError(t, os.Unsetenv(key))

	require.Equal(t, "v1", chooseDirWithRootEnv("v1", nil, key, "v2"))

	require.NoError(t, os.Setenv(key, "v3"))
	require.Equal(t, "v3", chooseDirWithRootEnv("", nil, key, "v2"))

	require.NoError(t, os.Unsetenv(key))
	require.NoError(t, os.Setenv(eKeyRoot, "root"))
	require.Equal(t, filepath.Join("root", "v2"), chooseDirWithRootEnv("", nil, key, "v2"))

	e1 := NewAppRootEnv("v4")
	require.Equal(t, filepath.Join("v4", "v2"), chooseDirWithRootEnv("", e1, key, "v2"))
}

func TestAppRootEnv(t *testing.T) {
	t.Run("init with value", func(t *testing.T) {
		e1 := NewAppRootEnv("demo")
		require.Equal(t, "demo", e1.RootDir())
		t.Run("twice", func(t *testing.T) {
			defer func() {
				require.NotNil(t, recover())
			}()
			e1.SetRootDir("test")
		})
	})

	t.Run("init with empty", func(t *testing.T) {
		e1 := NewAppRootEnv("")
		require.NoError(t, os.Setenv(eKeyRoot, "test"))
		defer func() {
			require.NoError(t, os.Unsetenv(eKeyRoot))
		}()
		require.Equal(t, "test", e1.RootDir())
		e1.SetRootDir("v2")
		require.Equal(t, "v2", e1.RootDir())
	})
}
