##### time.Sleep

##### mutex

##### chan

#### epoll
每个fd的封装类型都需要组合internal/poll.FD类型，然后由后者去做io  
对于进程继承来的标准fd(0,1,2)，先包装一下，然后internal/poll.FD依然使用阻塞方式去io  
对于自己打开的fd，也是先包装一下，然后internal/poll.FD统一使用异步复用方式进行io，golang使用协程把它模拟成阻塞方式  
```
listener/conn-----net.(*netFD)-------
                                     |-------internal/poll.(*FD)--------runtime跨平台支持poll
                  os.(*file) --------       (支持blocking和poll)
```

协程这边  
对fd进行io操作返回EAGAIN（读空或写满），需要调用netpollblock把自己阻塞起来  
在把自己阻塞之前先判断一下ioready，指不定恰好就绪，这样就不用阻塞了  
没什么希望的话就把pd.rg或者pd.wg值cas成wati，然后调用gopark把自己阻塞起来等待sysmon线程的唤醒  
唤醒后做的第一件事就是把pd.rg或者pd.wg的值xchg成0  
```
// returns true if IO is ready, or false if timedout or closed
// waitio - wait only for completed IO, ignore errors
func netpollblock(pd *pollDesc, mode int32, waitio bool) bool {
	gpp := &pd.rg
	if mode == 'w' {
		gpp = &pd.wg
	}

	// set the gpp semaphore to WAIT
	for {
		old := *gpp
		if old == pdReady {
			*gpp = 0
			return true
		}
		if old != 0 {
			throw("runtime: double wait")
		}
		if atomic.Casuintptr(gpp, 0, pdWait) {
			break
		}
	}

	// need to recheck error states after setting gpp to WAIT
	// this is necessary because runtime_pollUnblock/runtime_pollSetDeadline/deadlineimpl
	// do the opposite: store to closing/rd/wd, membarrier, load of rg/wg
	if waitio || netpollcheckerr(pd, mode) == 0 {
		gopark(netpollblockcommit, unsafe.Pointer(gpp), waitReasonIOWait, traceEvGoBlockNet, 5)
	}
	// be careful to not lose concurrent READY notification
	old := atomic.Xchguintptr(gpp, 0)
	if old > pdWait {
		throw("runtime: corrupted polldesc")
	}
	return old == pdReady
}
```

sysmon线程这边  
只要os触发了fd事件，就会调用netpollunblock函数去cas对应的pd.rg或者pd.wg把ioready的事件通知过去  
old可能是0，说明协程还在处理业务  
old可能是1，说明协程连上一次的ioready都没处理，可能是卡住了，那么也没必要做cas  
old可能是2，说明协程业务处理完了正准备block自己但被sysmon线程抢了先  
old可能是g，说明协程block自己等待ioready事件，sysmon会去唤醒它  
```
func netpollunblock(pd *pollDesc, mode int32, ioready bool) *g {
	gpp := &pd.rg
	if mode == 'w' {
		gpp = &pd.wg
	}

	for {
		old := *gpp
		if old == pdReady {
			return nil
		}
		if old == 0 && !ioready {
			// Only set READY for ioready. runtime_pollWait
			// will check for timeout/cancel before waiting.
			return nil
		}
		var new uintptr
		if ioready {
			new = pdReady
		}
		if atomic.Casuintptr(gpp, old, new) {
			if old == pdReady || old == pdWait {
				old = 0
			}
			return (*g)(unsafe.Pointer(old))
		}
	}
}
```