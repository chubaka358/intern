package list

// Node provides node interface
type Node interface {
	Set(next Node) Node
	GetValue() int
	GetNext() Node
}

// node implements list node
type node struct {
	val  int
	next Node
}

// GetValue returns node value
func (l *node) GetValue() int {
	return l.val
}

// GetNext returns node next
func (l *node) GetNext() Node {
	return l.next
}

// SetNextNode sets next node for node
func (l *node) Set(next Node) Node {
	l.next = next
	return l.next
}

// NewListNode returns new node
func NewListNode(val int) Node {
	return &node{
		val:  val,
		next: nil,
	}
}
