package main

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
