package strain

type Ints []int
type Lists [][]int
type Strings []string

func (list Ints) Keep(fx func(int) bool) Ints {
	var keep Ints
	for index := 0; index < len(list); index++ {
		if fx(list[index]) {
			keep = append(keep, list[index])
		}
	}
	return keep
}
func (list Ints) Discard(fx func(int) bool) Ints {
	var discard Ints
	for index := 0; index < len(list); index++ {
		if !fx(list[index]) {
			discard = append(discard, list[index])
		}
	}
	return discard
}
func (lists Lists) Keep(fx func([]int) bool) Lists {
	var keep [][]int
	for index := 0; index < len(lists); index++ {
		if fx(lists[index]) {
			keep = append(keep, lists[index])
		}
	}
	return keep
}
func (list Strings) Keep(fx func(string) bool) Strings {
	var keep Strings
	for index := 0; index < len(list); index++ {
		if fx(list[index]) {
			keep = append(keep, list[index])
		}
	}
	return keep
}
