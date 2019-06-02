package main

import (
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const (
	gMaxCount = 10000
)

func main() {
	runtime.GOMAXPROCS(3)
	var wg sync.WaitGroup
	var schedcnt int32
	st := time.Now()

	// 统计协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		// runtime.LockOSThread() // 意义？
		// defer runtime.UnlockOSThread()

		for {
			// if atomic.LoadInt32(&schedcnt) == gMaxCount {
			if schedcnt == gMaxCount {
				println("finish", "\t", time.Since(st).String())
				os.Exit(0)
			}
			println("schedcnt:", schedcnt, "\t", time.Since(st).String())
			time.Sleep(1 * time.Millisecond)
		}
	}()

	// 并发1万个协程做密集计算
	for i := 0; i < gMaxCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&schedcnt, 1)

			// // 密集计算场景1
			// count := 0
			// for {
			// 	// 常规累加，无栈检查
			// 	count++
			// 	runtime.Gosched()
			// }

			// 密集计算场景2
			// 每个协程运行10ms被抢占
			// 1万个协程，10000*10ms=100s
			// 3个p并发处理那么33s应该全部调度完毕
			// 然而实际却需要5m30s
			count := 0
			for {
				// hack累加，得嵌套两层才能骗取栈检查
				// 还有没有其它黑科技？
				func() {
					func() {
						count++
					}()
				}()
			}
		}()
	}
	wg.Wait()
	println("time cost:", time.Since(st).String())
}
