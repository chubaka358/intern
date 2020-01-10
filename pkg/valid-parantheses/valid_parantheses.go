package valid_parantheses

type Valider interface {
	IsValid(s string) bool
}

type valid struct {
}

// IsValid returns true if brackets sequence is valid
// else returns false
func (v *valid) IsValid(s string) bool {
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

func NewValid() Valider {
	return &valid{}
}
