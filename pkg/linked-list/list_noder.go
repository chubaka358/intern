package linked_list

// ListNoder provides listNode interface
type ListNoder interface {
	GetNext() ListNoder
	GetValue() int
	SetNext(next ListNoder)
}
