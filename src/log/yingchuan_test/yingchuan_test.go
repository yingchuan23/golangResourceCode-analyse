// Package log
// @Author chenyingchuan
// @Description //TODO
// @Date 22:48 2022/12/9
package yingchuan_test

import (
	"log"
	"os"
	"testing"
)

//ref https://blog.51cto.com/gotaly/1405754
func TestYingchuan(t *testing.T) {
	fileName := "xxx_debug.log"
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	debugLog := log.New(logFile, "[Debug]", log.Llongfile)
	debugLog.Println("A debug message here")
	debugLog.SetPrefix("[Info]")
	debugLog.Println("A Info Message here")
	debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	debugLog.Println("A different prefix")
}
