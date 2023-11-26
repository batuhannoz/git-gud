package main

// log(n)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{i, j}
		}
		m[num] = i
	}
	return []int{}
}