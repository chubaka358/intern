package linked_list

// ListNoder provides listNode interface
type ListNoder interface {
	SetNextNode(list ListNoder) ListNoder
	GetNext() ListNoder
}
