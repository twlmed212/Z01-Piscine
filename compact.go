package piscine

func Compact(ptr *[]string) int {
	start := 0
	for _, getSlice := range *ptr {
		if getSlice != "" {
			(*ptr)[start] = getSlice
			start++
		}
	}

	*ptr = (*ptr)[:start]
	return start
}
