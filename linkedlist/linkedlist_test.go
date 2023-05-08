package main

import (
	"errors"
	"testing"
)

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (l *LinkedList) First() *Node {
	return l.head
}
func (l *LinkedList) Last() *Node {
	return l.tail
}
func (l *LinkedList) Push(value int) {
	node := &Node{value: value}
	if l.head == nil { // head가 없을 경우 : head를 추가시키고 head의 previous를 nil로 만든다.
		l.head = node
		l.head.previous = nil
	} else {
		node.previous = l.tail
		l.tail.next = node // 리스트의 맨 마지막에 현재 노드를 추가
	}
	l.length++
	l.tail = node
}
func (l *LinkedList) Length() int {
	n := l.First()
	if n == nil {
		return 0
	}
	i := 1
	for {
		n = n.Next()
		if n == nil {
			break
		}
		i++
	}
	println("진짜 length", i)
	return l.length
}

type Node struct {
	value    interface{}
	next     *Node
	previous *Node
}

func (n *Node) Previous() *Node {
	return n.previous
}

func (n *Node) Next() *Node {
	return n.next
}

func (l *LinkedList) Shift() *Node {
	head := l.head
	newHead := l.head.Next()
	l.head = newHead
	l.head.previous = nil
	l.length--
	return head
}
func (l *LinkedList) Pop() *Node {
	tail := l.tail
	tail.Previous().next = nil
	//l.tail.previous = nil
	l.tail = tail.Previous()
	l.length--
	return tail
}
func (l *LinkedList) Remove(index int) {
	n := l.First()
	i := 0
	for {
		if i == index {
			back := n.Next() // 지워야 할 노드의 뒤를 잡음
			if index == 0 {
				back.previous = nil
				l.head = back
			} else {
				front := n.Previous() // 지워야 할 노드의 앞을 잡음 지금 있는 인덱스 제외
				front.next = back
				back.previous = front
			}
			l.length--
			break
		}
		n = n.Next()
		if n == nil {
			break
		}
		i++
	}

}

// 우선순위 변경

func (l *LinkedList) SetPriority(cur, targetIndex int) error { // 순서를 변경하기
	n := l.First()
	i := 0
	var currentNode *Node
	var targetNode *Node
	println(cur, targetIndex)
	for {
		println(i, cur, targetIndex)
		if i == targetIndex {
			targetNode = n
			println("target index gettodage")
		} else if i == cur {
			println("currentNode index gettodage")
			currentNode = n
		}
		if currentNode != nil && targetNode != nil {
			println("둘다 찾아서 멈춤")
			break
		}

		n = n.Next()
		if n == nil {
			return errors.New("next가 nil이어서 멈춤")
		}
		i++
	}

	//1. 현재노드의 앞뒤를 연결
	fp := currentNode.Previous()
	fn := currentNode.Next()

	fp.next = fn
	fn.previous = fp

	currentNode.previous = nil
	currentNode.next = nil

	//2. 타겟 노드의 앞뒤 사이로 현재노드 연결
	tp := targetNode.Previous()
	tn := targetNode.Next()

	tp.next = currentNode
	tn.previous = currentNode
	currentNode.previous = tp
	currentNode.next = tn
	return nil
}
func TestDoubleLinkedList(test *testing.T) {
	l := &LinkedList{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	n := l.First()
	println("--------", l.Length())
	for {
		println(n.value.(int))
		n = n.Next()
		if n == nil {
			break
		}
	}
	println("--------", l.Length())
	head := l.Shift()
	println("head : ", head.value.(int))
	println("now head : ", l.head.value.(int))
	println("--------add", l.Length())
	l.Push(4)
	l.Push(5)
	l.Push(6)
	println("--------", l.Length())
	n = l.First()
	for {
		println(n.value.(int))
		n = n.Next()
		if n == nil {
			break
		}
	}
	println("--------remove", l.Length())
	l.Remove(0)
	println("--------", l.Length())
	n = l.First()
	for {
		println(n.value.(int))
		n = n.Next()
		if n == nil {
			break
		}
	}
	l.Remove(2)
	println("--------", l.Length())
	n = l.Last()
	for {
		println(n.value.(int))
		n = n.Previous()
		if n == nil {
			break
		}
	}
	a := l.Pop()
	println(a.value.(int))
	println("--------", l.Length())
	n = l.First()
	for {
		println(n.value.(int))
		n = n.Next()
		if n == nil {
			break
		}
	}

	println("맨마지막꺼 : ", l.tail.value.(int))
	println("--------", l.Length())
	l.Push(5)
	l.Push(6)
	l.Push(7)
	l.Push(8)
	l.Push(9)
	l.Push(10)
	n = l.First()
	for {
		println(n.value.(int))
		n = n.Next()
		if n == nil {
			break
		}
	}
	println("--------", l.Length())
	err := l.SetPriority(6, 2)
	if err != nil {
		println("흑흑 ", err.Error())
	}
	println("--------", l.Length())
	n = l.First()
	for {
		println(n.value.(int))
		n = n.Next()
		if n == nil {
			break
		}
	}
	println("--------", l.Length())
}
