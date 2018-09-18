package series

func All(n int, s string) []string {
	var elementsToRet int = len(s) - n + 1
	arr := make([]string, elementsToRet)
	for i := 0; i < elementsToRet; i++ {
		arr[i] = s[i : i+n]
	}
	return arr
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if len(s) < n {
		return "", false
	}
	return s[:n], true
}
