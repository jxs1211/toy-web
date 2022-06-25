package main

func twoSum(s []int, target int) []int {
	m := make(map[int]int, len(s))
	for i, v := range s {
		_diff := target - v
		if value, ok := m[_diff]; ok {
			return []int{value, i}
		}
		m[v] = i
	}
	return []int{}
}
