package main

import "fmt"

func longestPalindrome(s string) int {
	if len(s) == 1 {
		return 1
	}
	// find number of even letters
	// add 1
	counter := make(map[string]int)
	for _, ch := range s {
		counter[string(ch)]++
	}
	total, odd := 0, 0
	for _, val := range counter {
		if val%2 == 0 {
			total++
		} else {
			odd++
		}
	}
	if odd == 0 {
		return len(s)
	}
	return len(s) - odd + 1
}

func main() {
	fmt.Println(longestPalindrome("abccccdd") == 7)
	fmt.Println(longestPalindrome("a") == 1)

}
