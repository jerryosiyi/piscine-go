package piscine

func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
	// ALTERNATE WAY
	// *a, *b = *b, *a
}
