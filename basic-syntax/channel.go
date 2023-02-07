package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收
	fmt.Println(x, y, x+y)

	// 通道缓冲区(缓冲区大小为2)
	ch := make(chan int, 2)
	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2
	// 获取这两个数据
	fmt.Println(<-ch, <-ch)

	// 通过内置len函数可获取ch元素个数，通过cap函数可获取ch的缓存容量长度

	// 遍历通道与关闭通道
	ch <- 3
	ch <- 4

	/*
	关闭未初始化的channel(nil)会panic
	重复关闭同一channel会panic
	向以关闭channel发送消息会panic
	从已关闭channel读取数据，不会panic，若存在数据，则可以读出未被读取的消息，若已被读出，则获取的数据为零值，可以通过ok-idiom的方式，判断channel是否关闭
	channel的关闭操作，会产生广播消息，所有向channel读取消息的goroutine都会接受到消息
	================================================================
	有些 channel 并不需要主动去关闭，因为 GC 会判断 channel 的状态并对闲置 channel 进行回收。
	只有那些需要明确告知相关方 channel 已关闭的场景，才需要主动进行 channel 的关闭，比如 for range 。
	 */
	close(ch)


	//for {
	//	if v, ok := <-ch; ok {
	//		fmt.Printf("%d ", v)
	//	} else {
	//		break
	//	}
	//}
	for v := range ch {
		fmt.Println(v)
	}



	/*
	单向channel
	其中 chan<- int表示只写channel, <-chan int表示只读channel，此类函数/方法声明可防止channel滥用，在编译时可以检测出。
	 */


	/*
	有缓存的channel使用环形数组实现，当缓存未满时，向channel发送消息不会阻塞，当缓存满时，发送操作会阻塞，
	直到其他goroutine从channel中读取消息；同理，当channel中消息不为空时，读取消息不会阻塞，当channel为空时，读取操作会阻塞，直至其他goroutine向channel发送消息。
	*/

	// 在某些应用场景，工作goroutine一直处理事务，直到收到退出信号
	mch := make(chan struct{})
	quit := make(chan struct{})
	for {
		select {
		case <- mch:
			// 正常工作
			//work()
		case <- quit:
			// 退出前，处理收尾工作
			//doFinish()
			return
		}
	}
}
