package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	Myslices := make([]int, max-min)

	count := 0
	for i := range Myslices {
		Myslices[i] = min + count
		if Myslices[i] != max {
			count++
		}
	}
	return Myslices
}
