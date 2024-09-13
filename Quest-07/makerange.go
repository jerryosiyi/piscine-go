package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	length := max - min
	arr := make([]int, length)
	for i := min; i < max; i++ {
		index := i - min
		arr[index] = i
	}
	return arr
}
