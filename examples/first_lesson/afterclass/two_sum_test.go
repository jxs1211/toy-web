package main

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		s      []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"base", args{[]int{2, 5, 5, 11}, 10}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.s, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSentienlError(t *testing.T) {
	eof := errors.New("EOF")
	println(eof)
	fmt.Printf("iof error: %p\n", eof)
	println(io.EOF)
	fmt.Printf("io.EOF error: %p\n", io.EOF)
	if eof == io.EOF {
		fmt.Println("equal")
	}
}
