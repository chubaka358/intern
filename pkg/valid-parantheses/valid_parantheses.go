package valid_parantheses

type Stack []rune

// newStack creates new stack
func newStack() *Stack {
	return &Stack{}
}

// push adds element to the top of stack
func (s *Stack) push(c rune) {
	*s = append(*s, c)
}

// pop delete and returns element from the top of stack
func (s *Stack) pop() rune {
	l := len(*s)
	tmp := (*s)[l-1]
	*s = (*s)[:l-1]
	return tmp
}

// len returns stack length
func (s *Stack) len() int {
	return len(*s)
}

// IsValid returns true if brackets sequence is valid
// else returns false
func IsValid(s string) bool {
	matchingBrackets := map[rune]rune{'{': '}', '[': ']', '(': ')'}
	stack := newStack()
	for _, c := range s {
		switch c {
		case '{', '[', '(':
			stack.push(c)
		case '}', ']', ')':
			if stack.len() < 1 {
				return false
			}
			tmp := stack.pop()
			if c != matchingBrackets[tmp] {
				return false
			}
		}
	}
	if stack.len() > 0 {
		return false
	}
	return true
}
