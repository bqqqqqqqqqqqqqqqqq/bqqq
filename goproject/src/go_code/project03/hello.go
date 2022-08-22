package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func putNum(numChan chan int) {
	for i := 1; i < 100000; i++ {
		numChan <- i
	}
	close(numChan)
}

func primeNum(numChan chan int, resChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-numChan
		if !ok {
			break
		}
		//假设是素数
		flag = true
		//判断是否为素数
		for i := 2; i < int(math.Sqrt(float64(num)))+1; i++ {

			if num%i == 0 {
				//不是素数
				flag = false
				break
			}
		}
		if flag {
			resChan <- num
		}
	}
	fmt.Println("此次协程已结束任务")
	exitChan <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	numChan := make(chan int, 2000000)
	resChan := make(chan int, 2000000)
	exitChan := make(chan bool, 8)
	start := time.Now().UnixMilli()

	go putNum(numChan)
	for i := 0; i < runtime.NumCPU(); i++ {
		go primeNum(numChan, resChan, exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(resChan)

		end := time.Now().UnixMilli()
		fmt.Printf("\n此次耗时间为：%v毫秒\n\n", end-start)
	}()

	for {
		res, ok := <-resChan
		if !ok {
			break
		}
		fmt.Printf("素数=“%d", res)
	}

}
