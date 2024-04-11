package piscine

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return a
	}
	start := nbrs[0]

	if len(nbrs) == 1 {
		if start < 0 {
			start = len(a) + start
			if start < 0 {
				return a
			}
		}

		return a[start:]
	}
	end := nbrs[1]

	if len(nbrs) == 2 {
		if start < 0 {
			start = len(a) + start
			if start > len(a) {
				start = len(a) // 5
			}
			if start < 0 {
				Slice := []string{}
				return Slice
			}
		}
		if end < 0 {
			end = len(a) + end
			if end > len(a) {

				end = len(a) /// 5
			}

		}
		if end == 0 {
			return nil
		}
		if end > start && end > len(a) {
			return nil
		}
		if start > end && start > len(a) {
			if end < len(a) {
				return a[:end]
			}
			return nil
		}
		return a[start:end]
	} else {
		Slice := []string{}
		return Slice

	}
	/*
		lenA := len(a)
		var mySlice []string
		if len(nbrs) == 2 {
			if nbrs[1] == 0 {
				return mySlice
			}
			if nbrs[0] < 0 {
				nbrs[0] = -nbrs[0]
				nbrs[0] = lenA - nbrs[0]
			}
			if nbrs[1] < 0 {
				nbrs[1] = -nbrs[1]
				nbrs[1] = lenA - nbrs[1]
			}
			return a[nbrs[0]:nbrs[1]]

		} else if len(nbrs) == 1 {
			if nbrs[0] < 0 {
				nbrs[0] = -nbrs[0]
				nbrs[0] = lenA - nbrs[0]
			}
			return a[nbrs[0]:]
		} else {
			fmt.Print("tete")
			return mySlice

		}*/

}
