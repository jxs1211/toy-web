package main

import "testing"

func Test_fibnacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"base", args{0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibnacci(tt.args.n); got != tt.want {
				t.Errorf("fibnacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
