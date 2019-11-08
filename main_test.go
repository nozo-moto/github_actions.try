package main

import "testing"

func Test_hoge(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "hoge",
			want: "hoge",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hoge(); got != tt.want {
				t.Errorf("hoge() = %v, want %v", got, tt.want)
			}
		})
	}
}
