package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func writeDataFile(c1 chan int, c2 chan bool) {
	for i := 1; i <= 1000; i++ {
		c1 <- rand.Intn(1000) + 1
	}
	close(c1)
	filePath := "/home/bqqq/goproject/test.txt"
	file, err := os.OpenFile(filePath, os.O_CREATE|
		os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for {
		v, ok := <-c1
		if !ok {
			break
		}
		str := strconv.Itoa(v) + "\n"
		writer.WriteString(str)
	}
	writer.Flush()
	c2 <- true
	close(c2)
}

func sortData(c1 chan int, c2 chan bool) {
	intSlice := make([]int, 0)
	file, err := os.Open("/home/bqqq/" +
		"goproject/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err1 := reader.ReadString('\n')
		if err1 != nil {
			if err1 == io.EOF {
				break
			}
			log.Fatal(err1)
		}

		// 提取出的str字符串含有'\n',需要去除
		// 采用strings.Trim(s string,cutset string)string
		new_str := strings.Trim(str, "\n")
		//将字符串转化为int 类型
		// lock.Lock()
		a, _ := strconv.ParseInt(new_str, 10, 64)
		intSlice = append(intSlice, int(a))
		// lock.Unlock()
	}
	sort.Ints(intSlice)
	// 将排序好后的数据写入管道
	for _, v := range intSlice {
		c1 <- v
	}
	close(c1)

	// 把管道中的数据写入新文件
	filePath := "/home/bqqq/goproject/test2.txt"
	file1, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	writer := bufio.NewWriter(file1)

	for {
		v, ok := <-c1
		if !ok {
			break
		}
		str := strconv.Itoa(v) + "\n"
		writer.WriteString(str)
		// fmt.Println(str)
	}
	writer.Flush()
	c2 <- true
	close(c2)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	intChan1 := make(chan int, 1000)
	exitChan1 := make(chan bool, 1)
	intChan2 := make(chan int, 10000)
	exitChan2 := make(chan bool, 1)
	go writeDataFile(intChan1, exitChan1)
	for {
		if _, ok := <-exitChan1; ok {
			break
		}
	}

	go sortData(intChan2, exitChan2)
	for {
		if _, ok := <-exitChan2; ok {
			break
		}
	}
	fmt.Println("test...")

}
