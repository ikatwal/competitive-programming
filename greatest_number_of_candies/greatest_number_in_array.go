package main

// 1431. Kids With the Greatest Number of Candies
func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxVal := 0
	// find max value
	for _, c := range candies {
		if c > maxVal {
			maxVal = c
		}
	}
	res := make([]bool, len(candies))
	// compare greatest value
	for i, c := range candies {
		if c+extraCandies >= maxVal {
			res[i] = true
		}
	}

	return res

}
