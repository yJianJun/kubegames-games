package poker

import (
	"fmt"
	"sync"
)

func interview() {
	channelA := make(chan int)
	channelB := make(chan int)
	channelC := make(chan int)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(3)
	go A(channelA, channelB, &waitGroup)
	go B(channelB, channelC, &waitGroup)
	go C(channelC, &waitGroup)
	channelA <- 0
	waitGroup.Wait()
}

func C(c chan int, s *sync.WaitGroup) {
	<-c
	fmt.Println("---------------C--------------")
	s.Done()
}

func B(b chan int, c chan int, s *sync.WaitGroup) {
	<-b
	fmt.Println("---------------B----------------")
	s.Done()
	c <- 2
}

func A(a chan int, b chan int, s *sync.WaitGroup) {
	<-a
	fmt.Println("---------------A-----------------")
	s.Done()
	b <- 1
}
