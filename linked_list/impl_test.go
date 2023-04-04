package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	t.Run("Test if new linked list is getting generated", func(t *testing.T) {
		list := LinkedList{}
		l := NewLinkedList()
		assert.Equal(t, list, l)
	})

	t.Run("Test push node", func(t *testing.T) {
		l1 := NewLinkedList()
		l1.Head = &Node{
			Data: 10,
		}
		l1.Length = 1
		l2 := NewLinkedList()
		l2.PushNode(10)
		assert.Equal(t, l1, l2)
		l2.PushNode(23)
		l2.PushNode(11)
		assert.Equal(t, 3, l2.Length)
	})

	t.Run("Test pop node", func(t *testing.T) {
		l1 := NewLinkedList()
		l1.PushNode(10)
		l1.PushNode(20)
		l1.PushNode(53)
		err := l1.PopNode()
		assert.Nil(t, err)
		assert.Equal(t, 2, l1.Length)
		l1 = NewLinkedList()
		err = l1.PopNode()
		assert.NotNil(t, err)

	})

	t.Run("Test insert node at index", func(t *testing.T) {
		l0 := NewLinkedList()
		l0.Length = 7
		l0.Head = &Node{
			Data: 23,
			Next: &Node{
				Data: 21,
				Next: &Node{
					Data: 10,
					Next: &Node{
						Data: 54,
						Next: &Node{
							Data: 2342,
							Next: &Node{
								Data: 43,
								Next: &Node{
									Data: 2,
								},
							},
						},
					},
				},
			},
		}
		l1 := NewLinkedList()
		err := l1.InsertAt(10, 2)
		assert.NotNil(t, err)

		err = l1.InsertAt(10, 0)
		assert.Nil(t, err)

		err = l1.InsertAt(21, 0)
		assert.Nil(t, err)

		l1.PushNode(43)
		assert.Equal(t, 3, l1.Length)

		err = l1.InsertAt(54, 2)
		assert.Nil(t, err)

		err = l1.InsertAt(2342, 3)
		assert.Nil(t, err)

		err = l1.InsertAt(23, 0)
		assert.Nil(t, err)

		err = l1.InsertAt(2, 6)
		assert.Nil(t, err)
		assert.Equal(t, 7, l1.Length)
		assert.Equal(t, l0, l1)
	})

	t.Run("Test delete node at index", func(t *testing.T) {
		l1 := NewLinkedList()
		err := l1.DeleteAt(0)
		assert.NotNil(t, err)

		err = l1.DeleteAt(-1)
		assert.NotNil(t, err)

		l1.PushNode(10)
		l1.PushNode(20)
		l1.PushNode(30)
		l1.PushNode(40)
		l1.PushNode(50)
		l1.PushNode(60)

		l1.DeleteAt(0)
		assert.Equal(t, 5, l1.Length)
		assert.Equal(t, " -> (20)  -> (30)  -> (40)  -> (50)  -> (60) ", l1.String())

		l1.DeleteAt(4)
		assert.Equal(t, 4, l1.Length)
		assert.Equal(t, " -> (20)  -> (30)  -> (40)  -> (50) ", l1.String())

		err = l1.DeleteAt(4)
		assert.NotNil(t, err)

		l1.DeleteAt(2)
		assert.Equal(t, 3, l1.Length)
		assert.Equal(t, " -> (20)  -> (30)  -> (50) ", l1.String())
	})
}
