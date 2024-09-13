package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	arr := []int{}
	for i := min; i < max; i++ {
		arr = append(arr, i)
	}
	return arr
}
