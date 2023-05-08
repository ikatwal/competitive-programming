package main

//605. Can Place Flowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i, pot := range flowerbed {
		if i == 0 {
			if i+1 != len(flowerbed) {
				if pot == 0 && flowerbed[i+1] == 0 {
					count++
					flowerbed[i] = 1
				}
			} else {
				if pot == 0 {
					count++
				}
			}
			continue
		}
		if i == len(flowerbed)-1 {
			if pot == 0 && flowerbed[i-1] == 0 {
				count++
				flowerbed[i] = 1
			}
			continue
		}
		if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 && pot == 0 {
			count++
			flowerbed[i] = 1
		}
	}

	return count >= n

}
