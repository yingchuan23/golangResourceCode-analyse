package yingchuan_os_test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"std/io/ioutil"
	"testing"
)

//ioUtil适合读取小文件
func TestIoUtil_yingchuan_read1(t *testing.T) {
	content, err := ioutil.ReadFile("./log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
}

//都是需要打开文件才能读取的，bufio 带缓存，可以读取大文件
func TestBufio_yingchuan_read2(t *testing.T) {
	file, err := os.Open("./log.txt")
	if err != nil {
		fmt.Println(file)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
}

//也可以根据自己定义的块（[]byte 切片长度）读取文件
func TestBufio2_yingchuan_read3(t *testing.T) {
	file, err := os.Open("./bufio.txt")
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(file)
	var recByteSlice []byte = make([]byte, 5)
	var content []byte
	for {
		_, err := reader.Read(recByteSlice)
		if err == io.EOF {
			break
		}
		content = append(content, recByteSlice...)
	}
	fmt.Println("文件内容已全部 append 到 content 切片")
	fmt.Println(string(content))
}
