package piscine

func Compact(ptr *[]string) int {
	var new []string
	for _, s := range *ptr {
		if s != "" {
			new = append(new, s)
		}
	}
	*ptr = new
	return len(new)
}
