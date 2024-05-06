package util

import "testing"

func TestHexToInt64(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "with prefix",
			args: args{
				hex: "0x12e3ccd",
			},
			want: 19807437,
		},
		{
			name: "without prefix",
			args: args{
				hex: "12e3ccd",
			},
			want: 19807437,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HexToInt64(tt.args.hex); got != tt.want {
				t.Errorf("HexToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64ToHex(t *testing.T) {
	type args struct {
		x int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				x: 19807437,
			},
			want: "0x12e3ccd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64ToHex(tt.args.x); got != tt.want {
				t.Errorf("Int64ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
