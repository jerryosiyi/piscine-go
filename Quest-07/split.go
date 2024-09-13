package piscine

func Split(s, sep string) []string {
	sepLen := len(sep)
	if sepLen == 0 {
		return []string{s}
	}
	start := 0 // keeps track of the start of current segment
	arr := []string{}
	for i := 0; i < len(s)-sepLen; {
		if s[i:i+sepLen] == sep {
			arr = append(arr, s[start:i])
			i += sepLen
			start = i
		} else {
			i++
		}
	}
	arr = append(arr, s[start:])
	return arr
}
