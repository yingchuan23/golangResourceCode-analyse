// Package yingchuan_runtime_test
// @Author chenyingchuan
// @Description //TODO
// @Date 23:46 2022/12/11
package yingchuan_runtime_test

import (
	"runtime"
	"testing"
)

// Caller 打印当前的行号
func TestRuntimeCaller1(t *testing.T) {
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		pcName := runtime.FuncForPC(pc).Name()
		t.Log("pc:", pc, "pcName:", pcName, "file:", file, "line:", line)
	}
}

// Caller 打印这个上层函数的行号
func TestRunTimeCaller2(t *testing.T) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		pcName := runtime.FuncForPC(pc).Name()
		t.Log("pc:", pc, "pcName:", pcName, "file:", file, "line:", line)
	}
}

// Caller 打印这个上上层函数的行号
func TestRuntimeCaller3(t *testing.T) {
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		pcName := runtime.FuncForPC(pc).Name()
		t.Log("pc:", pc, "pcName:", pcName, "file:", file, "line:", line)
	}
}
