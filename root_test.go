// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu
// Date: 2020/5/2

package fsenv

import (
	"os"
	"testing"
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
