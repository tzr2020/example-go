package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fileWrite2()
	fileRead3()
}

func fileWrite() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	defer file.Close()

	file.Write([]byte("Hello, world!"))
	// file.WriteString("Hello, world!")
}

func fileWrite2() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("Hello, world!\n") // 将数据写入缓存
	}
	writer.Flush() // 清空当前缓存，将缓存里的数据写入文件
}

func fileWrite3() {
	err := ioutil.WriteFile("./test.txt", []byte("Hello, world!"), 0666)
	if err != nil {
		fmt.Println("write file failed, err: ", err)
		return
	}
}

func fileRead() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file failed, err: ", err)
	}
	defer file.Close()

	var data []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			// 文件内容读取完毕
			break
		}
		if err != nil {
			fmt.Println("read file failed, err: ", err)
			return
		}
		data = append(data, tmp[:n]...)
	}
	fmt.Println(string(data))
}

func fileRead2() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file failed, err: ", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			// 文件内容读取完毕
			if len(line) != 0 {
				fmt.Println(line)
			}
			break
		}
		if err != nil {
			fmt.Println("read file failed, err: ", err)
			return
		}
		fmt.Print(line)
	}
}

func fileRead3() {
	data, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("read file failed, err: ", err)
		return
	}
	fmt.Print(string(data))
}
