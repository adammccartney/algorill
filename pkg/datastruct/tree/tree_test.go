package tree

import (
	"testing"
)

func Test_readProc(t *testing.T) {
	type args struct {
		miniproc string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test construction logic for tree",
			args: args{"(systemd) 1 0 0, (ModemManager) 2 1 0, (abrd-dbug) 3 1 0"},
			want: "(systemd) > (ModemManager) > (abrd-dbug)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readProc(tt.args.miniproc); got != tt.want {
				t.Errorf("readProc() = %v, want %v", got, tt.want)
			}
		})
	}
}
