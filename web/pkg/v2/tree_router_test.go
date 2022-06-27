package webv2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var test handlerFunc = func(c *Context) {}

func TestNewHandlerBasedOnTree(t *testing.T) {
	tests := []struct {
		name string
		want Handler
	}{
		// TODO: Add test cases.
		{"name", &HandlerBasedOnTree{root: &node{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandlerBasedOnTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandlerBasedOnTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlerBasedOnTree_Route(t *testing.T) {

	type fields struct {
		root *node
	}
	type args struct {
		method      string
		pattern     string
		handlerFunc handlerFunc
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"base", fields{&node{}}, args{http.MethodGet, "/test", test}, false},
		{"legal * pattern", fields{&node{}}, args{http.MethodGet, "/*", test}, false},
		{"illegal * pattern", fields{&node{}}, args{http.MethodGet, "abc*", test}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandlerBasedOnTree{
				root: tt.fields.root,
			}
			if err := h.Route(tt.args.method, tt.args.pattern, tt.args.handlerFunc); (err != nil) != tt.wantErr {
				t.Errorf("HandlerBasedOnTree.Route() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandlerBasedOnTree_ServeHTTP(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		c *Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandlerBasedOnTree{
				root: tt.fields.root,
			}
			h.ServeHTTP(tt.args.c)
		})
	}
}

func TestHandlerBasedOnTree_findMatchHandler(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   handlerFunc
		want1  bool
	}{
		// TODO: Add test cases.
		{"base", fields{&node{}}, args{"/test"}, test, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandlerBasedOnTree{
				root: tt.fields.root,
			}
			got, got1 := h.findMatchHandler(tt.args.path)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandlerBasedOnTree.findMatchHandler() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("HandlerBasedOnTree.findMatchHandler() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHandlerBasedOnTree_findMatchChild(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		root *node
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandlerBasedOnTree{
				root: tt.fields.root,
			}
			got, got1 := h.findMatchChild(tt.args.root, tt.args.path)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandlerBasedOnTree.findMatchChild() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("HandlerBasedOnTree.findMatchChild() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHandlerBasedOnTree_createSubTree(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		root      *node
		paths     []string
		handlerFn handlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandlerBasedOnTree{
				root: tt.fields.root,
			}
			h.createSubTree(tt.args.root, tt.args.paths, tt.args.handlerFn)
		})
	}
}

func TestHandlerBasedOnTree_validatePattern(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandlerBasedOnTree{
				root: tt.fields.root,
			}
			if err := h.validatePattern(tt.args.pattern); (err != nil) != tt.wantErr {
				t.Errorf("HandlerBasedOnTree.validatePattern() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newNode(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		// TODO: Add test cases.
		{"base", args{"/test"}, &node{
			path:     "/test",
			children: make([]*node, 2)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newNode(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny(t *testing.T) {
	// c := utf8.RuneCountInString("沈先捷")
	// fmt.Println(c)
	str := "沈先捷"
	b := []byte(str)
	fmt.Printf(" => bytes(hex): [% x]\n", b)
	b2, _ := json.Marshal(str)
	fmt.Printf("Marshal string => bytes(hex): [% x]\n", b2)
	var str2 interface{}
	_ = json.Unmarshal(b2, &str2)
	fmt.Printf("Unmarshal bytes => string bytes(hex): [% x]\n", str2)
	fmt.Printf(" => : %s\n", str2)
	fmt.Printf("is equal: %v\n", str == str2.(string)[1:]) // unmarhshal process is ok
	fmt.Printf("is equal: %v\n", str == str2.(string)[1:]) // something during unmarshal process
}
