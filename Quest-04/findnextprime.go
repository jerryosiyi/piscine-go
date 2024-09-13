package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	start := 0
	if nb%2 == 0 {
		start = nb + 1
	} else {
		start = nb
	}
	for i := start; ; i += 2 {
		if IsPrime(i) {
			return i
		}
	}
}
