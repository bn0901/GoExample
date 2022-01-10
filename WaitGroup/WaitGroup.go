/*
想要等待多个协程完成，我们可以使用 *wait group*。
题目目的: 并发进行5个工作。
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

// 每个协程都会运行该函数。 注意，WaitGroup 必须通过指针传递给函数。
func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done() // return 时，通知 WaitGroup，当前协程的工作已经完成

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second) // 睡眠一秒钟，以此来模拟耗时的任务。
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup  // WaitGroup 用于等待该函数启动的所有协程

	// 启动几个协程，并为其递增 WaitGroup 的计数器
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// 阻塞，直到 WaitGroup 计数器恢复为 0； 即所有协程的工作都已经完成
	wg.Wait()

	// results:
	//Worker 2 starting
	//Worker 1 starting
	//Worker 3 starting
	//Worker 5 starting
	//Worker 4 starting
	//Worker 3 done
	//Worker 1 done
	//Worker 2 done
	//Worker 5 done
	//Worker 4 done
	fmt.Println("End")
}