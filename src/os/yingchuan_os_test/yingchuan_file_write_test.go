// Package yingchuan_os
// @Author chenyingchuan
// @Description //TODO
// @Date 23:53 2022/12/10
package yingchuan_os_test

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileWrite1(t *testing.T) {
	file, err := os.OpenFile("writeTest.txt", os.O_WRONLY|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var logWriteSlice []byte = []byte("111111122222233333444141414214124212112")
	file.Write(logWriteSlice)
	//使用writeString 直接传输字符串
	file.WriteString("WriteString...\n")
	file.WriteString("WriteString...\n")
}

//使用ioutil -- 写小的方式
func TestIoUtilWrite(t *testing.T) {
	ioutil.WriteFile("writeTest.txt", []byte("ioutil write things"), 0666)
}

func TestBufioWrite(t *testing.T) {
	file, err := os.OpenFile("writeTest_bufio.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		err2 := file.Close()
		if err2 != nil {
			panic(err2)
		}
	}()
	writer := bufio.NewWriter(file)
	writer.WriteString("abcdefg\n")
	writer.Flush()
}

//只是Fprintf 输出到文件 他和printf 一样 只是输出不一样
func TestFprintf(t *testing.T) {
	file, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}
	serviceName := "yingchuan"
	FprintfTest1 := "%v writeTest for Fprintf Func\n"
	fmt.Fprintf(file, FprintfTest1, serviceName)

	fprintTest2 := []string{"111111", "222222", "3333333"}
	for _, value := range fprintTest2 {
		fmt.Fprintf(file, value)
	}
	defer file.Close()
}
