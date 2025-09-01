package study_case

import (
	"fmt"
	"sync"
)

// 互斥锁
var x2 int64
var wg2 sync.WaitGroup
var lock sync.Mutex // 1

func add2() {
	for i := 0; i < 5000; i++ {
		lock.Lock() //2 加锁
		x2 = x2 + 1
		lock.Unlock() //3 解锁
	}
	wg2.Done()
}
func SyncMutexCase() {
	wg2.Add(2)
	go add2()
	go add2()
	wg2.Wait()
	fmt.Println(x2)
}
