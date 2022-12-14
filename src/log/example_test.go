// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log_test

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

//TODO 学习使用 目前不知道下面的干啥
func ExampleLogger() {
	fmt.Println("看下打印")
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	logger.Print("Hello, log file!")
	fmt.Print(&buf)
	// Output:
	// logger: example_test.go:19: Hello, log file!
}

func TestLogger(t *testing.T) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	logger.Print("Hello, log file!")
	fmt.Print(&buf)
}

func ExampleLogger_Output() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "INFO: ", log.Lshortfile)

		infof = func(info string) {
			logger.Output(2, info)
		}
	)

	infof("Hello world")

	fmt.Print(&buf)
	// Output:
	// INFO: example_test.go:36: Hello world
}

func TestOutput_test(t *testing.T) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "INFO: ", log.Lshortfile)

		infof = func(info string) {
			logger.Output(2, info)
		}
	)
	infof("Hello world!!!")
	fmt.Print(&buf)
}
