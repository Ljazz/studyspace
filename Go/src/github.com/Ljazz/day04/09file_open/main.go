package main

import (
	"fmt"
	"os"
)

// 打开文件

func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	// 关闭文件
	defer fileObj.Close()
	// 读文件
	// var tmp = make([]byte, 128) // 指定读的长度
	var tmp [128]byte
	n, err := fileObj.Read(tmp[:])
	if err != nil {
		fmt.Printf("read from file failed, err:%v", err)
		return
	}
	fmt.Printf("读了%d个字节\n", n)
	fmt.Println(string(tmp[:n]))
}
