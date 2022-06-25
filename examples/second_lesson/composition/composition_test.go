package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestConcrete1_SayHello(t *testing.T) {
	type fields struct {
		Base Base
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{"base", fields{Base{"shen"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Concrete1{
				Base: tt.fields.Base,
			}
			c.SayHello()
		})
	}
}

func TestBase_SayHello(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Base{
				Name: tt.fields.Name,
			}
			b.SayHello()
		})
	}
}
