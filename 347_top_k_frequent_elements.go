package main


func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)   
	for _, num := range nums {
		m[num]++
	}
	
	f := make([][]int, 1 + len(nums))
	
	for key, val := range m {
		f[val] = append(f[val], key)
	}
	
	res := make([]int, 0, k)
	
	for i := len(f) - 1; i >= 0; i-- {
		if f[i] != nil {
			res = append(res, f[i]...)
			if len(res) >= k {
				res = res[:k]
				break
			}
		}
	}
	return res
}