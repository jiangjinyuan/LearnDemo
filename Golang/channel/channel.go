package main

import (
	"fmt"
)

// ReadClosedNoBufferChannel 接收已关闭非缓冲 channel
func ReadClosedNoBufferChannel() {
	// 非缓冲
	ch := make(chan int)

	close(ch)

	_, ok := <-ch
	if !ok {
		fmt.Println("channel closed, data invalid.")
	}
}

// ReadClosedBufferChannel 接收已关闭缓冲 channel
func ReadClosedBufferChannel() {
	ch := make(chan int, 5)
	ch <- 19
	close(ch)
	v, ok := <-ch
	if ok {
		fmt.Printf("received: %v\n", v)
	}
	v, ok = <-ch
	if !ok {
		fmt.Println("channel closed, data invalid.")
	}
}

func main() {
	ReadClosedNoBufferChannel()
	ReadClosedBufferChannel()
}
