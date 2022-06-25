package main

import "testing"

func Test_printNumWith2(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"base", args{0.222}, "0.22"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printNumWith2(tt.args.f); got != tt.want {
				t.Errorf("printNumWith2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printBytes(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"base", args{[]byte("Hello world")}, "Hello world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printBytes(tt.args.data); got != tt.want {
				t.Errorf("printBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
