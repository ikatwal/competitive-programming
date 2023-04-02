// leetcode 1480

package main

import "fmt"

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
