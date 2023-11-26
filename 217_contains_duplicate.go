package main

func containsDuplicate(nums []int) bool {
    if len(nums) <= 1 {
        return false
    }
    
    m := make(map[int]struct{})
    
    for _, num := range nums {
        if _, k := m[num]; k {
            return true
        }
        m[num] = struct{}{}
    }
    return false
}