package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type SElemType int

var Success = 1
var Fail = 0

func main() {
	// (1)
	// dbl := initDbl()
	// isE := dbl.isEmpty()
	// fmt.Println(isE)
	// dbl.push(0, SElemType(3))
	// fmt.Println(dbl)
	// isE = dbl.isEmpty()
	// fmt.Println(isE)
	// (2)
	// res := isS("sas")
	// fmt.Println(res)
	// (3)
	// aL := []int{1, 3, 5, -1, 7, 9, 10, -1, 9, -1}
	// pushStack(aL)
	// (4)
	// b := infix2ToPostfix("1+2*3+(4*5+6)*7")
	// fmt.Println(b)
	// a := calculate("123*+45*6+7*+")
	// fmt.Println(a)
	// (5)
	// fmt.Println(isValidSq([]string{"I", "O", "I", "I", "O", "I", "O", "O"}))
	// fmt.Println(isValidSq([]string{"I", "O", "O", "I", "O", "I", "I", "O"}))
	// fmt.Println(isValidSq([]string{"I", "I", "I", "O", "I", "O", "I", "O"}))
	// fmt.Println(isValidSq([]string{"I", "I", "I", "O", "O", "I", "O", "O"}))
	initQueue()
}

// (1)将编号为0和1的两个栈存放于一个数组空间V[m]中，栈底分别处千数组的两端。当第0号栈的栈顶指针top[O]等于-1时该栈为空;
// 当第1号栈的栈顶指针top[1]等于m时，该栈为空。
// 两个栈均从两端向中间增长(见图 3.17)。试编写双栈初始化，判断栈空、栈满、进栈和出栈等算法的函数。双栈数据结构的定义如下:
// typedef struct
//
//	int top[2J, bot[2); SElemType *V;
//	int m;
//
// ) DblStack;
const m = 9

type DblStack struct {
	top [2]int
	bot [2]int
	V   *[m]SElemType
	n   int
}

func initDbl() DblStack {
	Slist := [m]SElemType{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t := [2]int{-1, m}
	b := [2]int{-1, m}
	dbl := DblStack{t, b, &Slist, m}
	fmt.Println(dbl)
	return dbl
}

func (dbl DblStack) isEmpty() bool {
	if dbl.top[0] == -1 && dbl.top[1] == m {
		return true
	} else {
		return false
	}
}

func (dbl DblStack) isFull() bool {
	if dbl.top[0]+1 == dbl.top[1] {
		return true
	} else {
		return false
	}
}

func (dbl *DblStack) push(i int, e SElemType) {
	if dbl.top[0]+1 == dbl.top[1] {
		panic("stack is full!")
	}
	fmt.Println(dbl.top, dbl.V)

	if i == 0 {
		dbl.V[dbl.top[0]+1] = e
		dbl.top[0]++
	} else {
		dbl.V[dbl.top[0]-1] = e
		dbl.top[1]--
	}

}

func (dbl *DblStack) pop(i int, e SElemType) {
	if i == 0 && dbl.top[0] == -1 || (i == 1 && dbl.top[1] == m) {
		panic("stack is empty!")
	}

	if i == 0 {
		e = dbl.V[dbl.top[0]+1]
		dbl.top[0]--
	} else {
		e = dbl.V[dbl.top[0]-1]
		dbl.top[1]++
	}
}

// (2)回文是指正读反读均相同的字符序列，如 "abba" 和 "abdba" 均是回文，但 "good" 不 是回文。试写一个算法判定给定的字符序列是否为回文。(提示:将一半字符入栈)

type Stack struct {
	data []string
	top  int
	bot  int
}

func isS(e string) bool {
	eL := strings.Split(e, "")
	length := len(eL)
	mid := length / 2
	sD := make([]string, mid)
	s := Stack{data: sD, top: 0, bot: 0}
	if length == 0 {
		return true
	}

	for i := 0; i < mid; i++ {
		s.data[i] = eL[i]
	}
	for i := 0; i < mid; i++ {
		if s.data[i] != eL[length-i-1] {
			return false
		}
	}
	return true
}

// (3) 设从键盘输入一整数的序列:a1, a2, a3.., an,试编写算法实现:用栈结构存储输入的整数，当ai!=-1时，将ai进栈;当ai==-1时，输出栈顶整数并出栈。算法应对异常情况(入栈满等)给出相应的信息。
type Stack3 struct {
	data []int
	top  int
	bot  int
}

func (s *Stack3) push(e, n int) {
	if s.top == n {
		panic("stack is full")
	}
	s.data[s.top] = e
	s.top += 1
}

func (s *Stack3) pop(e int) {
	if s.top == 0 {
		panic("stack is empty")
	}
	e = s.data[s.top-1]
	s.top -= 1
}

func pushStack(l []int) {
	n := len(l)
	d := make([]int, n)
	s := Stack3{d, 0, 0}
	for _, value := range l {
		if value != -1 {
			s.push(value, n)
		} else {
			s.pop(value)
		}
	}
	fmt.Println(s.data, s)
}

// (4)从键盘上输入一个后缀表达式， 试编写算法计算表达式的值。规定:逆波兰表达式的长度不超过一行，以 "$"作为输入结束，操作数之间用空格分隔，操作符只可能有+、—、*、/四种 运算。例如: 23434 + 2*$。
type NumStack struct {
	data []string
	top  int
	bot  int
}

// 该解和题目描述有差异，直接输入一个后缀表达式字符串。然后输出计算值
func (s *NumStack) push(e string) {
	s.data[s.top] = e
	s.top += 1
}

func (s *NumStack) pop() string {
	e := s.data[s.top-1]
	s.top -= 1
	return e
}
func (s *NumStack) getTop() string {
	return s.data[s.top-1]
}
func (s *NumStack) isEmpty() bool {
	if s.top == 0 {
		return true
	}
	return false
}
func calculate(postfix string) int {
	d := make([]string, 100)
	stack := NumStack{d, 0, 0}
	fixLen := len(postfix)
	for i := 0; i < fixLen; i++ {
		nextChar := string(postfix[i])
		// 数字：直接压栈
		if unicode.IsDigit(rune(postfix[i])) {
			stack.push(nextChar)
		} else {
			// 操作符：取出两个数字计算值，再将结果压栈
			num1, _ := strconv.Atoi(stack.pop())
			num2, _ := strconv.Atoi(stack.pop())
			switch nextChar {
			case "+":
				stack.push(strconv.Itoa(num1 + num2))
			case "-":
				stack.push(strconv.Itoa(num1 - num2))
			case "*":
				stack.push(strconv.Itoa(num1 * num2))
			case "/":
				stack.push(strconv.Itoa(num1 / num2))
			}
		}
	}
	result, _ := strconv.Atoi(stack.getTop())
	fmt.Println(stack.data, stack.top)
	return result
}

// 中缀表达式转后缀表达式
// 转换过程
// 从左到右逐个字符遍历中缀表达式，输出的字符序列即是后缀表达式：
// 遇到数字直接输出
// 遇到运算符则判断：
// 栈顶运算符优先级更低则入栈，更高或相等则直接输出
// 栈为空、栈顶是(直接入栈
// 运算符是)则将栈顶运算符全部弹出，直到遇见)
// 中缀表达式遍历完毕，运算符栈不为空则全部弹出，依次追加到输出
// 中缀表达式转后缀表达式
func infix2ToPostfix(exp string) string {
	d := make([]string, 100)
	stack := NumStack{d, 0, 0}
	postfix := ""
	expLen := len(exp)
	// 遍历整个表达式
	for i := 0; i < expLen; i++ {
		char := string(exp[i])
		switch char {
		case " ":
			continue
		case "(":
			// 左括号直接入栈
			stack.push("(")
		case ")":
			// 右括号则弹出元素直到遇到左括号
			for !stack.isEmpty() {
				preChar := stack.getTop()
				if preChar == "(" {
					stack.pop() // 弹出 "("
					break
				}
				postfix += preChar
				stack.pop()
			}
		// 数字则直接输出
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			j := i
			digit := ""
			for ; j < expLen && unicode.IsDigit(rune(exp[j])); j++ {
				digit += string(exp[j])
			}
			postfix += digit
			i = j - 1 // i 向前跨越一个整数，由于执行了一步多余的 j++，需要减 1
		default:
			// 操作符：遇到高优先级的运算符，不断弹出，直到遇见更低优先级运算符
			for !stack.isEmpty() {
				top := stack.getTop()
				if top == "(" || isLower(top, char) {
					break
				}
				postfix += top
				stack.pop()
			}
			// 低优先级的运算符入栈
			stack.push(char)
		}
	}
	// 栈不空则全部输出
	for !stack.isEmpty() {
		postfix += stack.pop()
	}
	return postfix
}

// 比较运算符栈栈顶 top 和新运算符 newTop 的优先级高低
func isLower(top string, newTop string) bool {
	// 注意 a + b + c 的后缀表达式是 ab + c +，不是 abc + +
	switch top {
	case "+", "-":
		if newTop == "*" || newTop == "/" {
			return true
		}
	case "(":
		return true
	}
	return false
}

// (5)假设以I和O分别表示入栈和出栈操作。栈的初态和终态均为空，入栈和出栈的操作序列可表示为仅由I和O组成的序列，称可以操作的序列为合法序列，否则称为非法序列。
// 1)下面所示的序列中哪些是合法的?
// A. IOIIOIOO B. 10010110 C. IIIOIOIO D. IIIOOIOO
// 2)通过对1)的分析，写出一个算法，判定所给的操作序列是否合法。若合法，返回true, 否则返回false(假定被判定的操作序列巳存入一维数组中)。
func isValidSq(s []string) bool {
	in, out := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case "I":
			in += 1
		case "O":
			out += 1
		}
		if out > in {
			return false // 出栈次数必须小于入栈次数，否则无法出栈
		}
	}
	if in == out { //终态为空，那么入栈和出栈次数必须一样
		return true
	} else {
		return false
	}

}

// (6) 假设以带头结点的循环链表表示队列，并且只设一个指针指向队尾元素结点 (注意:不设头指针)， 试编写相应的置空队列、 判断队列是否为空、 入队和出队等算法。
type LinkNode struct {
	next *LinkNode
	data int
}

func initQueue() {
	a := []int{1, 2, 3, 4, 5, 7}
	hNode := LinkNode{nil, 0}
	lp := &hNode
	lp.next = &hNode
	for i := 0; i < len(a); i++ {
		lp = lp.in(a[i])
	}
	fmt.Println(lp.String())
}

func (l *LinkNode) Clear() {
	if l.next == l {
		return
	}
	var h, p *LinkNode
	h = l.next
	p = h.next
	for ; p != l; p = p.next {
		h.next = p.next
	}
	l = h
	l.next = l
}

func (l *LinkNode) isEmpty() bool {
	if l.next == l {
		return true
	} else {
		return false
	}
}
func (l *LinkNode) String() []int {
	h := l.next
	if h == l {
		return []int{}
	}
	h = h.next
	var i []int
	for ; h != l.next; h = h.next {
		i = append(i, h.data)
	}
	return i
}

func (l *LinkNode) in(e int) *LinkNode {
	node := LinkNode{l.next, e} // 队尾进，且下一个节点指向头结点
	l.next = &node              // 尾指针指向的节点的next指向新创建的节点
	l = &node                   // 把尾指针指向新创建的节点
	return l
}

func (l *LinkNode) out() *LinkNode {
	h := l.next // 取到头结点
	if h != l { // 为空的判断，不为空则把头结点指向下一个节点
		h.next = h.next.next // 队头出元素
	}
	return l
}

// (7) 假设以数组Q[m]存放循环队列中的元素，同时设置一个标志ta$, 以tag= 0 和tag= 1 来区别在队头指针 (front) 和队尾指针 (rear) 相等时，队列状态为 “空” 还是 “满 "。试编写与 此结构相应的插入 (enqueue) 和删除 (dequeue) 算法。 (8)如果允许在循环队列的两端都可以进行插入和删除操作。要求:
// 心 写出循环队列的类型定义;
// @写出 “从队尾删除“ 和 “从队头插入" 的算法。
// 85
//  .
// I 数据结构(C语言版)(第2版) l
// (9) 已知Ackermann 函数定义如下:
// n+l 当m=O时
// Ack(m,n) =�Ack(m -1, 1) 当m "'F-0,n=O时
// Ack(m -1,Ack(m,n -1))当m -ct:-0,n-ct:- 0时
// 心写出计算Ack(m, n)的递归算法，并根据此算法给出Ack(2, 1)的计算过程。 @写出计算Ack(m, n)的非递归算法。 (10)已知f为单链表的表头指针，链表中存储的都是整型数据，试写出实现下列运算的递归
// 算法:
// 心求链表中的最大整数; @求链表的结点个数; @求所有整数的平均值。
