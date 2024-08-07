// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/5/28

package fsenv

import (
	"os"
	"testing"

	"github.com/fsgo/fst"
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
				"k11": "a=11",
				"k12": "12#2",
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
			fst.Equal(t, tt.want, got)
		})
	}
}

func TestLoadEnvFile(t *testing.T) {
	t.Run("load success", func(t *testing.T) {
		fst.NoError(t, os.Unsetenv("k1"))
		defer func() {
			fst.NoError(t, os.Unsetenv("k1"))
		}()
		fst.Equal(t, "", os.Getenv("k1"))
		fst.NoError(t, LoadEnvFile("testdata/envfile/b.env"))
		fst.Equal(t, "v1", os.Getenv("k1"))
	})

	t.Run("file not exists", func(t *testing.T) {
		fst.Error(t, LoadEnvFile("testdata/envfile/not_exists.env"))
	})

	t.Run("bad format", func(t *testing.T) {
		fst.Error(t, LoadEnvFile("testdata/envfile/c_bad.env"))
	})
}

func TestMustLoadEnvFile(t *testing.T) {
	t.Run("load success", func(t *testing.T) {
		fst.NoError(t, os.Unsetenv("k1"))
		defer func() {
			fst.NoError(t, os.Unsetenv("k1"))
		}()
		fst.Equal(t, "", os.Getenv("k1"))
		MustLoadEnvFile("testdata/envfile/b.env")
		fst.Equal(t, "v1", os.Getenv("k1"))
	})

	t.Run("file not exists", func(t *testing.T) {
		defer func() {
			fst.NotNil(t, recover())
		}()
		MustLoadEnvFile("testdata/envfile/not_exists.env")
	})
}
