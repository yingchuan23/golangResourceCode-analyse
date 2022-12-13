// Package yingchuan_expand_knowledge
// @Author chenyingchuan
// @Description //闭包函数演示 优势在于可以直接使用函数里面的变量 ，不用在自己的进行生命
// @Date 15:38 2022/12/13
package yingchuan_expand_knowledge

import (
	"fmt"
	"testing"
)

func getSequence() func() int {
	var i int = 0
	return func() int {
		i++
		return i
	}
}

func TestBibaoFunc(t *testing.T) {
	sequence := getSequence()

	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	f := getSequence()
	fmt.Println(f())
	fmt.Println(f())
}
