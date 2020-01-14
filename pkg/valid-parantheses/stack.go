package valid_parantheses

// stacker provides stack interface
type stacker interface {
	push(c rune)
	pop() rune
	len() int
}

// stack implements stack structure
type stack []rune

// push adds element to the top of stack
func (s *stack) push(c rune) {
	*s = append(*s, c)
}

// pop delete and returns element from the top of stack
func (s *stack) pop() rune {
	l := len(*s)
	tmp := (*s)[l-1]
	*s = (*s)[:l-1]
	return tmp
}

// len returns stack length
func (s *stack) len() int {
	return len(*s)
}

// newStack creates new stack
func newStack() *stack {
	return &stack{}
}
