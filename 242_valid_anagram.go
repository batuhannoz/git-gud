package main

// log(n) 
func isAnagram(s string, t string) bool {
	m := make(map[rune]int)
	
	for _, char := range s {
		if _, ok := m[char]; ok {
			m[char]++
		} else {
			m[char] = 1
		}
	}
	
	for _, char := range t {
		if _, ok := m[char]; !ok {
			return false
		}
		if m[char] == 0 {
			return false
		}
		m[char]--
	}
	
	for _, e := range m {
		if e != 0 {
			return false
		}
	}
	
	return true
}