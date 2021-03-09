package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var n int64
var wg sync.WaitGroup

func main() {
	//runtime.GOMAXPROCS(1) //单核
	runtime.GOMAXPROCS(2) //多核
	wg.Add(10000)
	for i:=0;i<10000;i++{
		go add()
	}
	wg.Wait()
	fmt.Println("累加结果:",n)
}
func add() {
	for i := 0; i < 100; i++ {
		//n++
		atomic.AddInt64(&n, 1)
		//time.Sleep(1)
	}
	wg.Done()
}