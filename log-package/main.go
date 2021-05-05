package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	// 设置日志格式
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	// 设置日志前缀
	log.SetPrefix("[file]")

	logFile, err := os.OpenFile("./test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err: ", err)
		return
	}
	// 设置日志输出位置
	log.SetOutput(logFile)
}

func main() {
	log.Println("this a log")

	// 创建logger
	logger := log.New(os.Stdout, "[stdout]", log.Lmicroseconds|log.Lshortfile)
	logger.Println("this a log")
}
