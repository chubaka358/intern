package linked_list

// Cycler provides cycle interface
type Cycler interface {
	HasCycle(head ListNoder) bool
}
