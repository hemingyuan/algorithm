package base

import (
	"fmt"
	"testing"
)

func display(ll *LinkList) {
	var fn = func(node *LinkListNode) {
		fmt.Printf("node=%v ", node)
	}
	fmt.Println("visit start, length=", ll.length)
	ll.Visit(fn)
	fmt.Println("\nvisit end, length=", ll.length)
}

func TestLinkList(t *testing.T) {
	a := CreateLinkList()
	fmt.Println("a.length", a.length)
	a.Insert(1)
	display(a)
	a.Insert(2)
	a.Insert(3)
	display(a)
	a.InsertAfter(23, a.Find(2))
	a.InsertAfter(233, a.Find(2))
	display(a)
	a.DeleteByValue(233)
	display(a)
	a.Delete()
	display(a)
	a.Delete()
	display(a)
	fmt.Println("a has loop = ", a.HasLoop())
	b := CreateLinkList()
	b.Insert(9)
	b.Insert(6)
	b.Insert(5)
	b.Insert(2)

	c := CreateLinkList()
	c.Insert(8)
	c.Insert(7)
	c.Insert(4)
	c.Insert(1)

	d := MergeSortedLinkList(b, c)
	display(d)
}
