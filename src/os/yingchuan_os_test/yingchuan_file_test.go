// Package yingchuan_os_test
// @Author chenyingchuan
// @Description //TODO
// @Date 23:39 2022/12/9
package yingchuan_os_test

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestFile_yingchuan_open1(t *testing.T) {
	file, err := os.Open("./log.text")
	defer file.Close()
	if err != nil {
		fmt.Println("open file failed!,err: ", err)
		return
	}
}

func TestFile_yingchuan_open2(t *testing.T) {
	file, err := os.OpenFile("./log/txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
}

// 知道大小的情况
func TestFile_yingchuan_read1(t *testing.T) {
	file, err := os.Open("./log.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	var receiveByteSliceBuf []byte = make([]byte, 2048)
	n, err := file.Read(receiveByteSliceBuf)
	fmt.Printf("以供读了 %v 个字节\n", n)
	fmt.Println(receiveByteSliceBuf)
}

//不知道大小的情况
func TestFile_yingchuan_read2(t *testing.T) {
	file, err := os.Open("./log.txt")
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	for {
		var receiveByteSliceBuf = make([]byte, 3)
		_, err := file.Read(receiveByteSliceBuf)
		if err == io.EOF {
			break
		}
		fmt.Print(string(receiveByteSliceBuf))
	}
	fmt.Printf("\n")
	fmt.Println("读取文件结束。。。")
}

//输出到content
func TestFile_yingchuan_read3(t *testing.T) {
	file, err := os.Open("./log.txt")
	if err != nil {
		fmt.Println("open file failed ! ,err: ", err)
		return
	}
	//定义一个接受内容的content
	var content []byte
	for {
		var receiveByteSliceBuf = make([]byte, 3)
		_, err := file.Read(receiveByteSliceBuf)
		if err == io.EOF {
			break
		} else if err != nil {
			println("读取文件错误", err)
		}

		content = append(content, receiveByteSliceBuf...)
	}
	fmt.Printf("读取文件 %v 结束。。。 文件内容如下：\n", file.Name())
	fmt.Println(string(content))
}
