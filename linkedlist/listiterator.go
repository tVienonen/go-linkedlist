package linkedlist

// ListIterator can be used to traverse through a List that implements the Lister interface
type ListIterator struct {
	linkedList *Lister
	pos        *Node
}

// Next method returns the next item in the list
func (l *ListIterator) Next() *Node {
	node := l.pos
	if l.pos != nil {
		l.pos = l.pos.next
	}
	return node
}

// NewListIterator creates an iterator for a list
func NewListIterator(lister Lister) Iterator {
	return &ListIterator{
		linkedList: &lister,
		pos:        lister.Head()}
}
