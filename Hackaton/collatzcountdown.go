package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	n := start
	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			steps++
		} else {
			n = 3*n + 1
			steps++
		}
	}
	return steps
}
