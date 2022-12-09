// Package log
// @Author chenyingchuan
// @Description //编码是从上往下编写的，下面的代码更新可能会对上面的代码出现error
// @Date 22:48 2022/12/9
package yingchuan_test

import (
	"fmt"
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

func init() {
	log.SetPrefix("ServiceTest 服务日志：")
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	logFile, err := os.OpenFile("servicetest.log ", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.SetOutput(logFile)
	} else {
		fmt.Println("ServiceTest 日志文件创建/读写错误：", err)
	}
}

//ref https://www.itsky.tech/2022/07/27/Golang-Log-%E6%A0%87%E5%87%86%E5%BA%93-%E6%96%87%E4%BB%B6%E6%93%8D%E4%BD%9C/
func TestWwwitskytech(t *testing.T) {
	log.Println("this line is info log")
	//log.Panic("this line is panic error log")
	//log.Panicln("this line is panic error log")
	//log.Fatal("this line is fatal log")
	//log.Fatalln("this line is fatal log")

	fmt.Println(log.Flags())

	//2、重新设定 log.Flags log.Llongfile 是打印绝对路径，
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println(log.Flags())
	log.Println("this log.Flags is 13")
	log.Println("====================+")

	//3. prefix
	fmt.Println(log.Prefix())
	log.SetPrefix("RDS Service status ====> ")
	log.Println("[info] starting...")

	//4. log.SetOutput 配置日志的输出位置
	log.SetOutput(os.Stdout)
	log.Println("设置日志的打印的位置")

	//5、文件实现了 io.writer 接口，放一个文件对象
	file01, err := os.Create("./log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(file01)
	log.Println("当前日志的打印为值在log.txt 中。。。")

	//6、用os.Create 每次都是创建一个 io.Writer 接口对象
	file02, err := os.OpenFile("./log_openFile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(file02)
	log.Println("用 os.openFile 函数 打开文件并写入日志")
}
