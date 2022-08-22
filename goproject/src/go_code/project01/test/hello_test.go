package test

import "testing"

func TestCal(t *testing.T) {
	numChan := make(chan int, 2000)         //创建一个管道 存放我们需要计算的数值
	resChan := make(chan map[int]int, 2000) //创建一个管道 存放我们计算完成的结果
	exit := make(chan int, 8)               //创建一个管道 用户判断所有协程对于resChan写入操作是否完成计算
	cal(numChan, resChan, exit)

}
