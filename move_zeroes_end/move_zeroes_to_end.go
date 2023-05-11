package main

// 283. Move Zeroes

func moveZeroes(nums []int) {
	curr := 0
	for i, n := range nums {
		if n == 0 {
			continue
		}
		nums[i] = 0
		nums[curr] = n
		curr++
	}
}
