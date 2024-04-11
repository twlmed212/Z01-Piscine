package piscine

func Abort(a, b, c, d, e int) int {
	mySlice := []int{}
	get := append(mySlice, a, b, c, d, e)
	for i := 0; i < len(get); i++ {
		for j := 0; j < len(get); j++ {
			if get[i] > get[j] {
				junkvar := get[i]
				get[i] = get[j]
				get[j] = junkvar
			}
		}
	}
	return get[2]
}
