package poker

import (
	"fmt"
	"time"
)

// A首先被a阻塞，A()结束后关闭b，使b可读
func A(x, y chan int) {
	<-x
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(y)
}

// B首先被a阻塞，B()结束后关闭b，使b可读
func B(y, z chan int) {
	<-y
	fmt.Println("B()!")
	close(z)
}

// C首先被a阻塞
func C(z chan int) {
	<-z
	fmt.Println("C()!")
}

func interview() {
	x := make(chan int)
	y := make(chan int)
	z := make(chan int)

	//上1个执行结束 开启下1个trigger
	go A(x, y)
	go B(y, z)
	go C(z)

	// 关闭x，让x可读  关闭的channel中可以不断的读取零值
	close(x)
	time.Sleep(3 * time.Second)
}
