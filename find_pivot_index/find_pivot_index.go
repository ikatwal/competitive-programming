package main

import "fmt"

// Number 724
/**
Given an array of integers nums, calculate the pivot index of this array.
The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.
If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. This also applies to the right edge of the array.
Return the leftmost pivot index. If no such index exists, return -1.
*/

func pivotIndex(nums []int) int {
	// find sum of nums
	total := 0
	for _, n := range nums {
		total += n
	}
	// find pivot by comparing the total - sumsoFar - currentNumber
	sum := 0
	for pivot, n := range nums {
		// compare sum
		if sum == (total - sum - n) {
			return pivot
		}
		// add n to sum
		sum += n
	}
	return -1
}

func main() {
	input := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(input))
}
