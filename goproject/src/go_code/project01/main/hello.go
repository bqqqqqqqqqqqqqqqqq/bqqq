package main

import (
	"fmt"
	"runtime"
)

func cal(c chan int, resChan chan map[int]int, exit chan int) {
	for {
		k, ok := <-c
		if !ok {
			break
		}
		var resNum int
		for i := 1; i <= k; i++ {
			resNum += i
		}
		mymap := map[int]int{k: resNum}
		resChan <- mymap
	}
	exit <- 1
}

func main() {
	cores := runtime.NumCPU()
	fmt.Println(cores)
	runtime.GOMAXPROCS(cores - 1)
	numChan := make(chan int, 2000)
	resChan := make(chan map[int]int, 2000)
	exit := make(chan int, 8)

	go func() {
		for i := 0; i < 2000; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	for i := 0; i < 8; i++ {
		go cal(numChan, resChan, exit)

	}

	for i := 0; i < 8; i++ {
		<-exit
	}
	close(resChan)
	for v := range resChan {
		fmt.Println(v)

	}
}
