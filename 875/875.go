package koko_eating_bananas

func MinEatingSpeed(piles []int, h int) int {
	// Koko at least eats 1 banana per hour
	l, r := 1, 1

	// Find the max value from piles so that we could know max eating speed that Koko might need
	for _, p := range piles {
		r = max(r, p)
	}

	// Use binary search to get best eating speed to match guard's hours of leaving
	for l < r {
		mid := (l + r) / 2
		hours := 0

		// Iterate over the piles and calculate hours.
		// We increase the hours by ceil(pile / middle)
		for _, p := range piles {
			hours += (p + mid - 1) / mid
		}

		// Check if middle is a workable speed, and cut the search space by half
		if hours <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}

	// Once the left and right boundaries coincide, we find the target value,
	// that is, the minimum workable eating speed.
	return l
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
