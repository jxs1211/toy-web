package main

import "fmt"

// 输出两位小数
func printNumWith2(f float64) string {
	return fmt.Sprintf("%.2f", f)
}

func printBytes(data []byte) string {
	return string(data)
}
