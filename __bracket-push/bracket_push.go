package brackets

type stack struct {
	top    int
	array  []byte
	length int
}

func newStack() *stack {
	s := new(stack)
	s.top = -1
	s.array = make([]byte, 8)
	s.length = 32
	return s
}

func (stack *stack) Empty() bool {
	return stack.top == -1
}

func (stack *stack) Push(element byte) {
	if stack.top >= stack.length {
		stack.array = append(stack.array, make([]byte, 8)...)
	}
	stack.array[stack.top+1] = element
	stack.top++
}

func (stack *stack) Pop() (byte, bool) {
	if stack.Empty() {
		return 0, false
	}
	b := stack.array[stack.top]
	stack.top--
	return b, true
}

var startClose = map[byte]byte{
	123: 125,
	91:  93,
	40:  41,
}

var starters = map[byte]bool{
	123: true,
	91:  true,
	40:  true,
}
var closers = map[byte]bool{
	125: true,
	93:  true,
	41:  true,
}

func Bracket(s string) (bool, error) {

	stack := newStack()
	for i := range s {
		if starters[s[i]] {
			stack.Push(s[i])
		}
		if closers[s[i]] {
			lastWas, notEmpty := stack.Pop()
			if !notEmpty || startClose[lastWas] != s[i] {
				return false, nil
			}
		}
	}

	return stack.Empty(), nil
}
