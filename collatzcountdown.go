package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	count := 0
	if start == 1 {
		return count
	}
	if start%2 == 1 {
		// odd number
		count = 1 + CollatzCountdown(start*3+1)
	}
	if start%2 == 0 {
		// even Number
		count = 1 + CollatzCountdown(start/2)
	}
	return count
}
