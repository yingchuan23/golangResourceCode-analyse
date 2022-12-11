// Package yingchuan_sync_test
// @Author chenyingchuan
// @Description //TODO
// @Date 00:43 2022/12/11
package yingchuan_sync_test

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync"
	"testing"
)

//写着里面的函数不是完全的懂 为啥close(waitChan)
func TestMutliWriter(t *testing.T) {
	var writeChan chan interface{} = make(chan interface{})
	var isDone chan interface{} = make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go WriterRole(&wg, writeChan)
	}
	file, err := os.OpenFile("goroutine_writeTest.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	go FileReceiver(writeChan, file, isDone)

	wg.Wait()
	close(writeChan)

	<-isDone
}

// 写随机数函数
func WriterRole(wg *sync.WaitGroup, writeChan chan interface{}) {
	writeChan <- rand.Intn(100)
	defer wg.Done()
}

// 写入函数，循环取 writeChan 信道中的数据，直到 writeChan 信道关闭
func FileReceiver(writeChan chan interface{}, file io.Writer, isDone chan interface{}) {
	for res := range writeChan {
		fmt.Fprintln(file, res)
	}
	isDone <- "Done"
}
