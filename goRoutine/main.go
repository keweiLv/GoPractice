package main

import (
	"fmt"
	"sync"
	"time"
)

func mygo(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("In goroutine %s\n", name)
		time.Sleep(time.Millisecond * 1000)
	}
}

func main1() {
	go mygo("协程1号")
	go mygo("协程2号")
	time.Sleep(time.Minute)
}

func main2() {
	pipline := make(chan int, 10)
	fmt.Printf("信道可缓冲 %d 个数据\n", cap(pipline))
	pipline <- 1
	fmt.Printf("信道中当前有 %d 个数据", len(pipline))
}

func main3() {
	done := make(chan bool)
	fmt.Printf("%d\n", cap(done))
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	<-done
}

func worker(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("worker %d: %d\n", x, i)
	}
}

func main4() {
	var wg sync.WaitGroup

	wg.Add(2)
	go worker(1, &wg)
	go worker(2, &wg)

	wg.Wait()
}

func main() {
	lock := &sync.RWMutex{}
	lock.Lock()

	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Printf("第 %d 个协程准备开始... \n", i)
			lock.RLock()
			fmt.Printf("第 %d 个协程获得读锁, sleep 1s 后，释放锁\n", i)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)
	}

	time.Sleep(time.Second * 2)

	fmt.Println("准备释放写锁，读锁不再阻塞")
	// 写锁一释放，读锁就自由了
	lock.Unlock()

	// 由于会等到读锁全部释放，才能获得写锁
	// 因为这里一定会在上面 4 个协程全部完成才能往下走
	lock.Lock()
	fmt.Println("程序退出...")
	lock.Unlock()
}
