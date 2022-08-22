package test

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
}
