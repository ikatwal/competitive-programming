package main

import "fmt"

// # Number 1768
func mergeAlternately(word1 string, word2 string) string {
	l1, l2 := 0, 0
	res := []rune{}
	for l1 < len(word1) && l2 < len(word2) {
		res = append(res, rune(word1[l1]), rune(word2[l2]))
		l1++
		l2++
	}
	if l1 < len(word1) {
		res = append(res, []rune(word1[l1:])...)
	} else if l2 < len(word2) {
		res = append(res, []rune(word2[l2:])...)
	}

	return string(res)
}

func main() {
	// abc pqrs
	// apbqcrs
	fmt.Println(mergeAlternately("abc", "pqrs") == "apbqcrs")
}
