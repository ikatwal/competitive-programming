package main

import "fmt"

// solved 704

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			hi = mid - 1
		} else if target > nums[mid] {
			lo = mid + 1
		}
	}

	return -1
}

func main() {
	in := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(in, 9) == 4)
	fmt.Println(search(in, 19) == -1)

}
