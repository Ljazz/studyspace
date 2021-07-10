package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		文件操作：
		1. 路径
			相对路径：relative
				相对于当前的工程
			绝对路径：absolute
				从根目录开始

			. 当前目录
			.. 上一层
		2. 创建文件夹，如果文件夹存在，创建失败
			os.MkDir()：创建一层
			os.MkDIrAll()：可以创建多层

		3. 创建文件，Create采用模式0666（任何人可读写， 不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
			os.Create()：创建文件

		4. 打开文件，让当前程序和指定的文件之间创建一个连接
			os.Open(filename)
			os.OpenFile(filename, mode perm)

		5. 关闭文件，程序和文件之间的连接断开
			file.Close()

		6. 删除文件或目录
			os.Remove()：删除文件和空目录
			os.RemoveAll()：删除所有
	*/
	// 1. 路径
	// fileName1 := "D:/Git_projects/studyspace/Go/src/github.com/Ljazz/day04/04/aa.txt"
	// fileName2 := "bb.txt"
	// fmt.Println(filepath.IsAbs(fileName1)) // true
	// fmt.Println(filepath.IsAbs(fileName2)) // false
	// fmt.Println(filepath.Abs(fileName1))   // D:\Git_projects\studyspace\Go\src\github.com\Ljazz\day04\04\aa.txt <nil>
	// fmt.Println(filepath.Abs(fileName2))   // D:\Git_projects\studyspace\Go\src\github.com\Ljazz\day04\04\bb.txt <nil>

	// fmt.Println("获取父目录：", path.Join(fileName1, ".."))

	// 2. 创建文件
	// err := os.Mkdir("D:/Git_projects/studyspace/Go/src/github.com/Ljazz/day04/04/aa/bb", os.ModePerm)
	// if err != nil {
	// 	fmt.Println("err: ", err)
	// 	return
	// }
	// fmt.Println("文件夹创建成功")
	// err := os.MkdirAll("D:/Git_projects/studyspace/Go/src/github.com/Ljazz/day04/04/aa/cc/dd/ee", os.ModePerm)
	// if err != nil {
	// 	fmt.Println("err: ", err)
	// 	return
	// }
	// fmt.Println("多层文件夹创建成功")

	// 3. 创建文件：Create采用模式0666（任何人可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
	// file1, err := os.Create("D:/Git_projects/studyspace/Go/src/github.com/Ljazz/day04/04/aa/ab.txt")
	// if err != nil {
	// 	fmt.Println("err: ", err)
	// 	return
	// }
	// fmt.Println(file1)

	// file2, err := os.Create("file2.txt") // 创建相对路径的文件，是以当前工程目录为参照
	// if err != nil {
	// 	fmt.Println("err: ", err)
	// 	return
	// }
	// fmt.Println(file2)

	// 4. 打开文件
	// file3, err := os.Open("aa.txt")
	// if err != nil {
	// 	fmt.Println("err: ", err)
	// 	return
	// }
	// fmt.Println(file3)

	/*
		第一个参数：文件名称
		第二个参数：文件的打开方式
			const (
				// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
				O_RDONLY int = syscall.O_RDONLY // open the file read-only.
				O_WRONLY int = syscall.O_WRONLY // open the file write-only.
				O_RDWR   int = syscall.O_RDWR   // open the file read-write.
				// The remaining values may be or'ed in to control behavior.
				O_APPEND int = syscall.O_APPEND // append data to the file when writing.
				O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
				O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
				O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
				O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
			)
		第三个参数：文件的权限：文件不存在创建文件，需要指定权限
	*/
	// file4, err := os.OpenFile("aa.txt", os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	// if err != nil {
	// 	fmt.Println("err: ", err)
	// 	return
	// }
	// fmt.Println(file4)

	// 5. 关闭文件
	// file4.Close()

	// 5. 删除文件或文件夹
	// err := os.Remove("bb.txt")
	// if err != nil {
	// 	fmt.Println("err: ", err)
	// 	return
	// }
	// fmt.Println("删除文件成功...")
	// 删除目录
	// err := os.Remove("D:/Git_projects/studyspace/Go/src/github.com/Ljazz/day04/04/aa/cc/dd/ee")
	// if err != nil {
	// 	fmt.Println("err: ", err)
	// 	return
	// }
	// fmt.Println("删除目录成功...")
}

func printFileInfo() {
	/*
		FileInfo：文件信息
			interface
				Name()，文件名
				Size()，文件大小，字节为单位
				IsDir()，是否是目录
				ModTime()，修改时间
				Mode()，权限
	*/
	fileInfo, err := os.Stat("./aa.txt")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Printf("文件名：%v\n", fileInfo.Name())
	fmt.Printf("文件大小：%v\n", fileInfo.Size())
	fmt.Printf("是否是目录：%v\n", fileInfo.IsDir())
	fmt.Printf("修改时间：%v\n", fileInfo.ModTime())
	fmt.Printf("文件权限：%v\n", fileInfo.Mode())
	fmt.Println(fileInfo.Sys())
}
