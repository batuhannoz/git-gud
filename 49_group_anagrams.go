package main

func groupAnagrams(strs []string) [][]string {
    m := map[[26]int][]string{}
    
    for _, s := range strs {
        k := [26]int{}
        
        for i := 0; i < len(s); i++ {
            k[s[i]-'a'] += 1
        }
        m[k] = append(m[k], s)
    }
    res := [][]string{}
    
    for _, v := range m {
        res = append(res, v)
    }
    return res
}