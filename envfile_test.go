// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParserEnvFile(t *testing.T) {
	type args struct {
		fn string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "a.env",
			args: args{
				fn: "testdata/envfile/a.env",
			},
			want: map[string]string{
				"k1":  "v1",
				"k2":  "v2",
				"k3":  " v3",
				"k4":  "v4 ",
				"k5":  "v5 ",
				"k6":  "v6",
				"k7":  "v7",
				"k8":  "",
				"k9":  "",
				"k10": " ",
			},
			wantErr: false,
		},
		{
			name: "b.env",
			args: args{
				fn: "testdata/envfile/b.env",
			},
			want: map[string]string{
				"k1": "v1",
			},
			wantErr: false,
		},
		{
			name: "a.env",
			args: args{
				fn: "testdata/envfile/not_exists.env",
			},
			wantErr: true,
		},
		{
			name: "c_bad.env",
			args: args{
				fn: "testdata/envfile/c_bad.env",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParserEnvFile(tt.args.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParserEnvFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestLoadEnvFile(t *testing.T) {
	t.Run("load success", func(t *testing.T) {
		require.NoError(t, os.Unsetenv("k1"))
		defer func() {
			require.NoError(t, os.Unsetenv("k1"))
		}()
		require.Equal(t, "", os.Getenv("k1"))
		require.NoError(t, LoadEnvFile("testdata/envfile/b.env"))
		require.Equal(t, "v1", os.Getenv("k1"))
	})

	t.Run("file not exists", func(t *testing.T) {
		require.Error(t, LoadEnvFile("testdata/envfile/not_exists.env"))
	})

	t.Run("bad format", func(t *testing.T) {
		require.Error(t, LoadEnvFile("testdata/envfile/c_bad.env"))
	})
}

func TestMustLoadEnvFile(t *testing.T) {
	t.Run("load success", func(t *testing.T) {
		require.NoError(t, os.Unsetenv("k1"))
		defer func() {
			require.NoError(t, os.Unsetenv("k1"))
		}()
		require.Equal(t, "", os.Getenv("k1"))
		MustLoadEnvFile("testdata/envfile/b.env")
		require.Equal(t, "v1", os.Getenv("k1"))
	})

	t.Run("file not exists", func(t *testing.T) {
		defer func() {
			require.NotNil(t, recover())
		}()
		MustLoadEnvFile("testdata/envfile/not_exists.env")
	})
}
