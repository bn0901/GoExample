/*
Go 中最主要的状态管理机制是依靠通道间的通信来完成的。
我们已经在工作池的例子中看到过，并且，还有一些其他的方法来管理状态。
这里，我们来看看如何使用 sync/atomic 包在多个协程中进行 _原子计数_。
 */
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)  //ops: 50000
}