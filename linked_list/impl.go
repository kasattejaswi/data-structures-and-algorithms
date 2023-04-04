package linked_list

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Head   *Node
	Length int
}

type Node struct {
	Data int
	Next *Node
}

func NewLinkedList() LinkedList {
	return LinkedList{
		Head:   nil,
		Length: 0,
	}
}

func (l *LinkedList) PushNode(data int) {
	if l.Head == nil {
		l.Head = &Node{
			Data: data,
			Next: nil,
		}
		l.Length += 1
	} else {
		p := l.Head
		for {
			if p.Next == nil {
				p.Next = &Node{
					Data: data,
				}
				break
			}
			p = p.Next
		}
		l.Length += 1
	}
}

func (l *LinkedList) PopNode() error {
	if l.Head == nil {
		return errors.New("no node to pop")
	}
	if l.Head.Next == nil {
		l.Head = nil
		l.Length -= 1
		return nil
	}
	p := l.Head
	for {
		if p.Next.Next == nil {
			p.Next = nil
			l.Length -= 1
			return nil
		}
		p = p.Next
	}
}

func (l *LinkedList) InsertAt(value int, index int) error {
	if index < 0 {
		return errors.New("index cannot be less than 0")
	}

	if index > l.Length {
		return errors.New("index exceeded linked list length")
	}

	if l.Head == nil && index == 0 {
		l.Head = &Node{
			Data: value,
		}
		l.Length += 1
		return nil
	}

	if index == 0 {
		n := Node{
			Data: value,
			Next: l.Head,
		}
		l.Head = &n
		l.Length += 1
		return nil
	}
	var i int = 0
	p := l.Head
	for {
		if index == l.Length && index == i+1 {
			p.Next = &Node{
				Data: value,
				Next: nil,
			}
			l.Length += 1
			return nil
		}
		if index == i+1 {
			p.Next = &Node{
				Data: value,
				Next: p.Next,
			}
			l.Length += 1
			return nil
		}
		i++
		p = p.Next
	}
}

func (l *LinkedList) DeleteAt(index int) error {
	if index < 0 {
		return errors.New("index cannot be less than 0")
	}
	if index == l.Length {
		return errors.New("index cannot be greater than list's length")
	}
	var i int
	p := l.Head
	if index == 0 {
		l.Head = l.Head.Next
		l.Length--
		return nil
	}
	for {
		if index == i+1 && index == l.Length-1 {
			p.Next = nil
			l.Length--
			return nil
		}
		if index == i+1 {
			p.Next = p.Next.Next
			l.Length--
			return nil
		}
		p = p.Next
		i++
	}
}

func (l LinkedList) String() string {
	var s string
	p := l.Head
	for {
		s += fmt.Sprintf(" -> (%v) ", p.Data)
		if p.Next == nil {
			break
		}
		p = p.Next
	}
	return s
}
