package main

import "time"

// func main() {
// 	c := make(chan bool)
// 	go func() {
// 		if v, ok := <-c; ok {
// 			fmt.Println(v)
// 		} else {
// 			fmt.Println("chan closed")
// 		}
// 	}()

// 	// 发送方关闭chan
// 	time.Sleep(1 * time.Second)
// 	close(c)
// 	time.Sleep(1 * time.Second)

// 	fmt.Println("exit")
// }

// func main() {
// 	c := make(chan bool)
// 	go func() {
// 		time.Sleep(2 * time.Second) // 等待协程起来
// 		close(c)
// 	}()
// 	c <- true

// 	fmt.Println("exit")
// }

// func main() {
// 	c := make(chan bool, 1)
// 	go func() {
// 		time.Sleep(2 * time.Second) // 等待协程起来
// 		close(c)
// 	}()
// 	c <- true
// 	c <- true

// 	fmt.Println("exit")
// }

// func main() {
// 	c := make(chan bool)
// 	close(c)
// 	c <- true //send on closed channel

// 	fmt.Println("exit")
// }

// func main() {
// 	i := 0
// 	go func() {
// 		print(i)
// 	}()

// 	time.Sleep(1 * time.Second)
// 	i++
// 	print(i)
// }
// 注释掉i++，变量i变成只读，main.func1得到i的副本
// 去掉go关键字，变量i在栈上，main.func1得到i的引用
// 保留go关键字，变量i在堆上，main.func1同样得到i的引用

// func main() {
// 	c := make(chan int, 5)
// 	var wg sync.WaitGroup
// 	// 多写
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			c <- i
// 			wg.Done()
// 		}(i)
// 	}

// 	wg.Wait()

// // 方式一，有明确写入数的，写完就close然后用for range读
// //       因为for range读一个没有关闭但已经空了的channle会deadlock
// // close(c)
// for v := range c {
// 	print(v)
// }

// 	// 方式二，写入数量不确定的就无法close，没close的就不能用for range，只能用for死循环
// 	for {
// 		select {
// 		case v := <-c:
// 			print(v)
// 		default:
// 			break
// 		}
// 	}
// }

// func main() {
// 	c := make(chan int, 5)
// 	for i := 0; i < 5; i++ {
// 		c <- i
// 	}
// 	go func() {
// 		c <- 5
// 	}()

// 	time.Sleep(time.Second)
// 	// close(c)
// 	for v := range c {
// 		print(v)
// 	}
// }
// 验证FIFO

func main() {
	for i := 0; i < 10; i++ {
		ch := make(chan int, i)
		go func() {
			cap := cap(ch)
			println(cap)
			close(ch)
		}()
	}
	time.Sleep(time.Second)
}
