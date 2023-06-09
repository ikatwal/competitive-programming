package main

import "fmt"

// Leetcode 392. Is Subsequence

func main() {
	s, t := "abc", "ahbgdc"
	fmt.Println(isSubsequence(s, t))
	s = "axc"
	fmt.Println(isSubsequence(s, t))
}

// returns True if s1 is a subsequence of s2 else False
func isSubsequence(s1, s2 string) bool {
	s1Index, s2Index := 0, 0
	// while both are left to traverse
	for s1Index < len(s1) && s2Index < len(s2) {
		// increment s1
		if s1[s1Index] == s2[s2Index] {
			s1Index += 1
		}
		s2Index += 1
	}
	return s1Index == len(s1)
}
