package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return nb
	}
	low, high := 1, nb
	for low <= high {
		mid := (low + high) / 2
		midSquared := mid * mid

		if midSquared == nb {
			return mid
		} else if midSquared < nb {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return 0
}
