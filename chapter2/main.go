package main

import (
	"fmt"
)

type ElemType int

type LNode struct {
	data ElemType
	next *LNode
}

type LinkList *LNode

func main() {
	// (1)
	// l1, l2 := initL()
	// Connect(l1, l2)
	// printList(l1)
	// (2)
	// l1, l2 := initL1()
	// l3 := Connect1(l1, l2)
	// printList(l3)
	// (3)
	// l1, l2 := initL()
	// Intersection1(l1, l2)
	// printList(l1)
	// (4)
	// l1, l2 := initL1()
	// Differ(l1, l2)
	// printList(l1)
	// (5)
	// l1 := initL5()
	// l2, l3 := divideList(l1)
	// printList(l2)
	// printList(l3)
	// (6)
	// l1 := initL5()
	// max := maxNode(l1)
	// fmt.Println(max)
	// (7)
	// l1 := initL5()
	// reverse(l1)
	// printList(l1)
	// (8)
	// l1 := initL8()
	// deleteList(l1, -4, 7)
	// printList(l1)
	// (9)
	// l9 := initL9()
	// change(*l9.next.next.next.next)
	// printDuList(l9)
	// (10)
	// delItem()
}

// (1)将两个递增的有序链表合并为一个递增的有序链表。要求结果链表仍使用原来两个链表的存储空间，不另外占用其他的存储空间。表中不允许有重复的数据。
func Connect(Ta, Tb LinkList) {
	var pa, pb, pc *LNode
	pc = Ta // pc指向合并后的链表当前节点
	pa = Ta.next
	pb = Tb.next                 // pa,pb指向两个原始链表的当前节点
	for pa != nil && pb != nil { // pa,pb指针在两个原始链表中移动，将较小的作为新链表的下一个节点
		if pa.data < pb.data {
			pc.next = pa
			pa = pa.next
		} else if pa.data == pb.data {
			pc.next = pa
			pa = pa.next
			pb = pb.next
		} else {
			pc.next = pb
			pb = pb.next
		}
		pc = pc.next
	}
	if pa != nil { // 由于是递增的链表，所以将还未移动到最后的链表直接并在新链表后面即可
		pc.next = pa
	} else {
		pc.next = pb
	}
	Tb = nil
	// Ta为合并后链表
}
func initL() (*LNode, *LNode) {
	a := [9]ElemType{3, 3, 4, 5, 6, 7, 8, 9, 12}
	b := [9]ElemType{1, 4, 6, 7, 8, 11, 14, 15, 16}
	l1 := initList(a)
	l2 := initList(b)
	return l1, l2
}

// (2)将两个非递减的有序链表合并为一个非递增的有序链表。要求结果链表仍使用原来两个链表的存储空间，不另外占用其他的存储空间。表允许有重复的数据。
// 非递减 -> 递增及相等 非递增 -> 递减及相等
// 使用头插法,但是节点还是从头开始遍历，相当于倒转了顺序
func Connect1(Ta, Tb LinkList) LinkList {
	var pa, pb, pc *LNode
	var Tc LinkList = Ta
	pa = Ta.next
	pb = Tb.next
	Tc.next = nil

	for pa != nil || pb != nil {
		if pa == nil {
			pc = pb
			pb = pb.next
		} else if pb == nil {
			pc = pa
			pa = pc.next
		} else if pa.data <= pb.data {
			pc = pa
			pa = pa.next
		} else {
			pc = pb
			pb = pb.next
		}

		pc.next = Tc.next
		Tc.next = pc

	}
	Tb = nil
	printList(Tc)
	return Tc
}

func initL1() (*LNode, *LNode) {
	a := [9]ElemType{1, 3, 4, 5, 5, 7, 8, 9, 12}
	b := [9]ElemType{3, 4, 6, 7, 7, 11, 14, 15, 16}
	l1 := initList(a)
	l2 := initList(b)
	return l1, l2
}

// (3)已知两个链表A和B分别表示两个集合，其元素递增排列。请设计一个算法，用于求出A与B的交集，并存放在A链表中。
func Intersection(Ta, Tb LinkList) { // n^2
	var pa, pb, pc *LNode
	pa = Ta.next
	pc = Ta
	for pa != nil {
		pb = Tb.next
		for pb != nil {
			if pa.data == pb.data {
				pc.next = pb
				pc = pc.next
				break
			}
			pb = pb.next
		}
		pa = pa.next
	}
	pc.next = nil
}
func Intersection1(Ta, Tb LinkList) { // n+n
	var pa, pb, pc *LNode
	pa = Ta.next
	pb = Tb.next
	pc = Ta
	for pa != nil && pb != nil {
		if pa.data == pb.data {
			pc.next = pa
			pa = pa.next
			pb = pb.next
			pc = pc.next
		} else if pa.data < pb.data {
			pa = pa.next
		} else {
			pb = pb.next
		}
	}
	pc.next = nil
}

// (4)巳知两个链表A和B分别表示两个集合，其元素递增排列。请设计算法求出两个集合A和B的差集(即仅由在A中出现而不在B中出现的元素所构成的集合)，并以同样的形式存储，同时返回该集合的元素个数。
func Differ(Ta, Tb LinkList) { // n+n
	var pa, pb, pc *LNode
	pa = Ta.next
	pb = Tb.next
	pc = Ta
	for pa != nil && pb != nil {
		if pa.data == pb.data {
			pa = pa.next
			pb = pb.next
		} else if pa.data < pb.data {
			pc.next = pa
			pa = pa.next
			pc = pc.next
		} else {
			pb = pb.next
		}
	}
	if pa != nil {
		pc.next = pa
	} else {
		pc.next = nil
	}
}

// (5)设计算法将一个带头结点的单链表A分解为两个具有相同结构的链表B和C,其中B表的结点为A表中值小于零的结点，而C表的结点为A表中值大于零的结点(链表A中的元素为非零整数，要求B、C表利用A表的结点)。
func initL5() *LNode {
	a := [9]ElemType{-1, 3, -4, 5, 5, 7, -8, 9, 12}
	l1 := initList(a)
	return l1
}
func divideList(Ta LinkList) (LinkList, LinkList) {
	var Tb, Tc LinkList
	var pa, pb, pc *LNode // 可以使用头插法，尾插法需要多两个指针
	Tb = Ta
	pb = Tb
	nodeC := LNode{0, nil}
	Tc = &nodeC
	pc = Tc
	pa = Ta.next
	for pa != nil {
		if pa.data > 0 {
			pc.next = pa
			pc = pc.next
		} else if pa.data < 0 {
			pb.next = pa
			pb = pb.next
		}
		pa = pa.next
	}
	pb.next = nil
	pc.next = nil
	return Tb, Tc
}

// (6)设计一个算法，通过一趟遍历确定长度为n的单链表中值最大的结点。
func maxNode(Ta LinkList) ElemType {
	var p *LNode
	p = Ta.next
	max := p.data
	if p.next == nil {
		return max
	}
	p = p.next
	for p != nil {
		if p.data > max {
			max = p.data
		}
		p = p.next
	}
	return max
}

// (7)设计一个算法，将链表中所有结点的链接方向“原地”逆转，即要求仅利用原表的存储空间，换句话说，要求算法的空间复杂度为0(1)。
func reverse(Ta LinkList) {
	var p, q *LNode // p记录当前遍历的节点
	p = Ta.next     //使用头插法，但是从头遍历原始链表，那么相当于倒转。
	Ta.next = nil   // 初始化新链表为空链表
	for p != nil {
		q = p.next
		p.next = Ta.next
		Ta.next = p
		p = q
	}
}

// (8)设计一个算法，删除递增有序链表中值大于mink且小于maxk:的所有元素(mink和maxk是给定的两个参数，其值可以和表中的元素相同，也可以不同)。
func initL8() *LNode {
	a := [9]ElemType{-7, -5, -4, 5, 6, 7, 8, 9, 12}
	l1 := initList(a)
	return l1
}
func deleteList(Ta LinkList, minK ElemType, maxK ElemType) {
	var fDelPrior, lDelNext, p *LNode
	p = Ta.next
	fDelPrior = p
	for p != nil {
		if p.data > minK {
			p = p.next
			break
		}
		fDelPrior = p
		p = p.next
	}
	for p != nil {
		if p.data < maxK {
			lDelNext = p.next
			break
		}
		p = p.next
	}
	if lDelNext != nil {
		fDelPrior.next = lDelNext
	} else {
		fDelPrior.next = nil
	}

}

// (9)巳知p指向双向循环链表中的一个结点，其结点结构为data、prior、next三个域，写出算法change(p), 交换p所指向的结点及其前驱结点的顺序。
func initL9() *DulNode {
	a := [9]ElemType{-7, -5, -4, 5, 6, 7, 8, 9, 12}
	l1 := initDuList(a)
	return l1
}

type DulNode struct {
	prior *DulNode
	data  ElemType
	next  *DulNode
}

func initDuList(s1 [9]ElemType) *DulNode {
	linkList := DulNode{nil, 9, nil}
	var q DuList
	q = &linkList
	for i := 0; i < 9; i++ {
		var p DulNode
		p.data = s1[i]
		q.next = &p
		p.prior = q
		q = &p
	}
	return &linkList
}
func change(p DulNode) { // 交换涉及四个结点，p的前驱，p的前驱的前驱，p的后继
	p.prior.next = p.next   // p的前驱的next指向p的后继
	p.next.prior = p.prior  // p的后继的prior指向p的前驱
	p.next = p.prior        // p的next指向p的前驱
	p.prior = p.prior.prior // p的prior指向p的前驱的前驱结点
	p.next.prior = &p       // 此时p.next已经指向了p的前驱，故前驱的prior指向p
	p.prior.next = &p       // 原来p的前驱的前驱的next指向p
}

type DuList *DulNode

func printDuList(s1 DuList) {
	var q DuList
	for s1 != nil {
		fmt.Println(s1.data)
		q = s1
		s1 = s1.next
	}
	fmt.Println("piror")
	for q != nil {
		fmt.Println(q.data)
		q = q.prior
	}
}

// (10)已知长度为n的线性表A采用顺序存储结构，请写一个时间复杂度为O(n)、空间复杂度为0(1)的算法，该算法可删除线性表中所有值为item的数据元素。
func delItem() {
	a := []int{1, 22, 99, 4, 5, 6, 7, 89, 99, 99, 99, 1, 23, 56}

	item := 99
	k := 0 // 记录等于item的元素个数
	for i := 0; i < len(a); i++ {
		if a[i] == item {
			k += 1
			continue // 等于item的元素丢弃
		}
		a[i-k] = a[i] // 前面有k个元素丢弃掉，则当前空位置为i-k，把原本i位置的非item值交换到该位置
	}
	length := len(a) - k
	a = a[:length]
	fmt.Println(a)
}

func initList(s1 [9]ElemType) *LNode {
	linkList := LNode{9, nil}
	for i := 9; i > 0; i-- {
		var p LNode
		p.data = s1[i-1]
		p.next = linkList.next
		linkList.next = &p

	}
	return &linkList
}

func printList(s1 LinkList) {
	for s1 != nil {
		fmt.Println(s1.data)
		s1 = s1.next
	}
}
