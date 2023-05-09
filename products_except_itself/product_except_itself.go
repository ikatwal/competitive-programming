package main

// 238. Product of Array Except Self

func productExceptSelf(nums []int) []int {
	// each numbers product is the product of left and product of right
	// we can think of it as i = (i-1)*(i+1)
	// so we keep track of sum of multiples at each i
	// so we traverse from left for each i (i-1)
	// and we traverse from right for each i (i+1)
	// the answer is their product
	// add 1 to left, as 0th element left multiple is 1
	// add 1 to end as nth element right multiple is 1
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	l := make([]int, 0)
	l = append(l, 1)
	r := make([]int, 0)
	r = append(r, 1)
	for i := 1; i < len(nums)-1; i++ {
		l = append(l, nums[i]*l[i-1])
	}
	rIndex := 0
	for j := len(nums) - 1; j >= 0; j-- {
		r = append(r, nums[j]*r[rIndex])
		rIndex++
	}
	answer := make([]int, len(nums)-2)
	rIndex = len(nums) - 2
	for k := 0; k < len(nums)-2; k++ {
		answer[k] = l[k] * r[rIndex]
		rIndex--
	}
	return answer
}
