package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var max int
	for i, r := range s {
		rs := fmt.Sprintf("%c", r)
		unique := make(map[string]bool)
		unique[rs] = true
		var count int = 1
		for _, c := range s[i+1:] {
			w := fmt.Sprintf("%c", c)
			if unique[w] {
				break
			}
			unique[w] = true
			count++
		}
		if count > max {
			max = count
		}
	}
	return max
}
