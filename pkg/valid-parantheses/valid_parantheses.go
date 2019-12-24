package valid_parantheses

type Stack []rune

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(c rune) {
	*s = append(*s, c)
}

func (s *Stack) Pop() rune {
	l := len(*s)
	tmp := (*s)[l-1]
	*s = (*s)[:l-1]
	return tmp
}

func (s *Stack) Len() int {
	return len(*s)
}

func IsValid(s string) bool {
	matchingBrackets := map[rune]rune{'{': '}', '[': ']', '(': ')'}
	stack := New()
	for _, c := range s {
		switch c {
		case '{', '[', '(':
			stack.Push(c)
		case '}', ']', ')':
			if stack.Len() < 1 {
				return false
			}
			tmp := stack.Pop()
			if c != matchingBrackets[tmp] {
				return false
			}
		}
	}
	if stack.Len() > 0 {
		return false
	}
	return true
}
