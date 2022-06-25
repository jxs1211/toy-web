package main

import "fmt"

type Parent struct {
}

func (p Parent) SayHello() string {
	return fmt.Sprintf("I am %s", p.Name())
}

func (p Parent) Name() string {
	return "Parent"
}

type Son struct {
	Parent
}

// func (s Son) SayHello() string {
// 	return fmt.Sprintf("I am %s", s.Name())
// }

// 定义了自己的 Name() 方法
func (s Son) Name() string {
	return "Son"
}
