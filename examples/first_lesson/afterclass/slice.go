package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4} // 直接初始化了 4 个元素的切片
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))

	s2 := make([]int, 3, 4) // 创建了一个包含三个元素，容量为4的切片
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	// s2 目前 [0, 0, 0], append（追加）一个元素，变成什么？
	s2 = append(s2, 7) // 后边添加一个元素，没有超出容量限制，不会发生扩容
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 8) // 后边添加了一个元素，触发扩容
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s3 := make([]int, 4) // 只传入一个参数，表示创建一个含有四个元素，容量也为四个元素的
	// 等价于 s3 := make([]int, 4, 4)
	fmt.Printf("s3: %v, len %d, cap: %d \n", s3, len(s3), cap(s3))

	// 按下标索引
	fmt.Printf("s3[2]: %d", s3[2])
	// 超出下标范围，直接崩溃
	// runtime error: index out of range [99] with length 4
	// fmt.Printf("s3[99]: %d", s3[99])

	// SubSlice()

	//shareArr()
}

func Add(s []int, index, value int) []int {
	new_s := make([]int, 0, len(s)+1)
	l_part, r_part := s[:index], s[index:]
	new_r_part := make([]int, 0, len(r_part))
	for i := 0; i < len(r_part); i++ {
		new_r_part = append(new_r_part, r_part[i])
	}
	l_part = append(l_part, value)
	new_s = append(new_s, l_part...)
	new_s = append(new_s, new_r_part...)
	return new_s
}

func Delete(s []int, index int) []int {
	return s
}
