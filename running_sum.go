package main

import "fmt"

// leetcode 1480
// Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
// Return the running sum of nums.
func runningSum(nums []int) []int {
	prev := 0
	for i := range nums {
		// add prev
		nums[i] = nums[i] + prev
		prev = nums[i]
	}
	return nums
}

func main() {
	input := []int{1, 2, 3, 4}
	fmt.Println(runningSum(input))
}
