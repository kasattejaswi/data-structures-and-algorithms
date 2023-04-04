package linked_list

// List interface defines a sequence specification
type List interface {
	PushNode(data int)
	PopNode() error
	InsertAt(value int, index int) error
	DeleteAt(index int) error
}
