package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
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

// 如何优雅的关闭 channel:
// 情形一：一个sender,一个receiver
// 情形二：一个sender,多个receiver
// 情形三：多个sender,一个receiver
// 情形四：多个sender,多个receiver

// 对于情形一、二，直接在sender端关闭；
// 对于情形三，增加一个传递关闭信号的 channel，receiver 通过信号 channel 下达关闭数据 channel 指令。senders 监听到关闭信号后，停止发送数据。

// 原则：不要在接收端关闭channel；当一个channel有多个发送方时，不要关闭channel；（不要重复关闭channel或者向一个关闭channel发送数据）

// CloseChannel01 关闭channel，针对情形三
func CloseChannel01() {
	rand.Seed(time.Now().UnixNano())

	const Max = 100000
	const NumSenders = 1000

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	// the receiver
	go func() {
		for value := range dataCh {
			if value == Max-1 {
				fmt.Println("send stop signal to senders.")
				close(stopCh)
				return
			}

			fmt.Println(value)
		}
	}()

	select {
	case <-time.After(time.Hour):
	}
	// 在 Go 语言中，对于一个 channel，如果最终没有任何 goroutine 引用它，不管 channel 有没有被关闭，最终都会被 gc 回收。
	// 所以，在这种情形下，所谓的优雅地关闭 channel 就是不关闭 channel，让 gc 代劳。
}

// CloseChannel02 关闭channel, 针对情形四
func CloseChannel02() {
	rand.Seed(time.Now().UnixNano())

	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	// It must be a buffered channel.
	toStop := make(chan string, NumReceivers+NumSenders)

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					toStop <- "sender#" + id
					return
				}

				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			for {
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						toStop <- "receiver#" + id
						return
					}

					fmt.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	select {
	case <-time.After(time.Hour):
	}
}

// happened before
// 关于 channel 的发送（send）、发送完成（send finished）、接收（receive）、接收完成（receive finished）的 happened-before 关系如下：
//	1、第 n 个 send 一定 happened before 第 n 个 receive finished，无论是缓冲型还是非缓冲型的 channel。
//	2、对于容量为 m 的缓冲型 channel，第 n 个 receive 一定 happened before 第 n+m 个 send finished。
//	3、对于非缓冲型的 channel，第 n 个 receive 一定 happened before 第 n 个 send finished。
//	4、channel close 一定 happened before receiver 得到通知。

func SendHappenedBeforeReceiveFinished() {
	var done = make(chan bool)
	var msg string

	go func() {
		msg = "Send HappenedBefore ReceiveFinished"
		done <- true
	}()

	<-done
	println(msg)
}

func ReceiveHappenedBeforeSendFinished() {
	var done = make(chan bool)
	var msg string

	go func() {
		msg = "Receive HappenedBefore SendFinished"
		<-done
	}()

	done <- true
	println(msg)
}

func main() {
	ReadClosedNoBufferChannel()
	ReadClosedBufferChannel()

	SendHappenedBeforeReceiveFinished()
	ReceiveHappenedBeforeSendFinished()
}
