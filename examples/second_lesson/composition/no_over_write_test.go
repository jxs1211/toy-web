package main

import (
	"testing"
)

func TestParent_Name(t *testing.T) {
	tests := []struct {
		name string
		p    Parent
		want string
	}{
		// TODO: Add test cases.
		{"base", Parent{}, "Parent"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parent{}
			if got := p.Name(); got != tt.want {
				t.Errorf("Parent.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSon_Name(t *testing.T) {
	type fields struct {
		son Son
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"base", fields{Son{}}, "Son"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Son{
				Parent: Parent{},
			}
			if got := s.Name(); got != tt.want {
				t.Errorf("Son.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParent_SayHello(t *testing.T) {
	tests := []struct {
		name string
		p    Parent
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parent{}
			if got := p.SayHello(); got != tt.want {
				t.Errorf("Parent.SayHello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSon_SayHello(t *testing.T) {
	tests := []struct {
		name string
		p    Son
		want string
	}{
		// TODO: Add test cases.
		{"base", Son{}, "I am Son"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Son{}
			if got := p.SayHello(); got != tt.want {
				t.Errorf("Son.SayHello() = %q, want %q", got, tt.want)
			}
		})
	}
}
