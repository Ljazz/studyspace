package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打开文件写内容
func writerDemo1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	defer fileObj.Close()
	// write
	fileObj.Write([]byte("zhoulin mengbi le"))
	// writeString
	fileObj.WriteString("周琳解释不了")
}

func writerDemo2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	defer fileObj.Close()
	// 创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello 天朝")
	wr.Flush()
}

func main() {
	writerDemo2()
}
