package listops

// IntList as
type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Reverse reverses the given list
func (list IntList) Reverse() IntList {
	reversedList := make(IntList, len(list))
	for index := 0; index < len(list); index++ {
		reversedList[index] = list[len(list)-1-index]
	}
	return reversedList
}

// Length returns the length of the list
func (list IntList) Length() int {
	return len(list)
}

// Append appends a list to the list
func (list IntList) Append(listToAppend IntList) IntList {
	newList := make(IntList, len(list)+len(listToAppend))
	for index := 0; index < len(list); index++ {
		newList[index] = list[index]
	}
	for index := 0; index < len(listToAppend); index++ {
		newList[len(list)+index] = listToAppend[index]
	}
	return newList
}

// Concat appends multiple lists to the list
func (list IntList) Concat(listsToAppend []IntList) IntList {
	totalLength := len(list)
	for _, v := range listsToAppend {
		totalLength += len(v)
	}
	newList := make(IntList, totalLength)
	for index := 0; index < len(list); index++ {
		newList[index] = list[index]
	}
	runningLength := len(list)
	for _, v := range listsToAppend {
		for index := 0; index < len(v); index++ {
			newList[runningLength+index] = v[index]
		}
		runningLength += len(v)
	}
	return newList
}

// Map returns a list with a function applied to each element
func (list IntList) Map(fx unaryFunc) IntList {
	newList := make(IntList, len(list))
	for index := 0; index < len(list); index++ {
		newList[index] = fx(list[index])
	}
	return newList
}

// Filter returns a list with elements satisfying a function
func (list IntList) Filter(fx predFunc) IntList {
	if len(list) == 0 {
		return IntList{}
	}
	var newList IntList
	for _, v := range list {
		if fx(v) {
			newList = newList.Append(IntList{v})
		}
	}
	return newList
}

// Foldr applies a function end to first
func (list IntList) Foldr(fx binFunc, start int) int {
	if len(list) == 0 {
		return start
	}
	result := fx(list[len(list)-1], start)
	for index := len(list) - 2; index > -1; index-- {
		result = fx(list[index], result)
	}
	return result
}

// Foldl applies a function first to end
func (list IntList) Foldl(fx binFunc, start int) int {
	if len(list) == 0 {
		return start
	}
	result := fx(start, list[0])
	for index := 1; index < len(list); index++ {
		result = fx(result, list[index])
	}
	return result
}
