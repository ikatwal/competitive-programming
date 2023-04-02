package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	return scan(s, t) && scan(t, s)
}

func scan(s, t string) bool {
	// ensure a 1 to 1 mapping for each character
	link := make(map[string]string)
	for i, ch := range s {
		if _, ok := link[string(ch)]; ok {
			if string(t[i]) != link[string(ch)] {
				return false
			}
		} else {
			link[string(ch)] = string(t[i])
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("a", "b"))
}
