package piscine

func DescendAppendRange(max, min int) []int {
	mySlice := []int{}

	if max <= min {
		return mySlice
	} else {
		for i := max; i > min; i-- {
			mySlice = append(mySlice, i)
		}
	}
	return mySlice
}

// The max must be included, and min must be excluded.
// If max is inferior than or equal to min, return an empty slice.
// make() is not allowed for this exercise.
