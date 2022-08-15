package main

import (
	"fmt"
	"unsafe"
)

const age = 12

func main() {
	helloWord()
	strAppent()
	varType()
	constVar()
	selDemo()
	forDemo()
}

// 打印hello word
func helloWord() {
	fmt.Println("hello word")
}

// str 使用 + 拼接
func strAppent() {
	fmt.Println("hello" + " " + "word")
}

func varType() {
	fmt.Println(age)
}

func constVar() {
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	fmt.Println(a, b, c)

}

func selDemo() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("received ", i1, "from c1\n")
	case c2 <- i2:
		fmt.Println("sent ", i2, "to c2\n")
	case i3, ok := <-c3:
		if ok {
			fmt.Println("received ", i3, "from c3\n")
		} else {
			fmt.Println("c3 is closed\n")
		}
	default:
		fmt.Println("no comunication\n")
	}
}

func forDemo() {
	var result int = 0
	const flag int = 10
	for i := 0; i < flag; i++ {
		result += i
	}
	fmt.Println(result)

}
