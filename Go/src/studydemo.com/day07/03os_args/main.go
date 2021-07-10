package main

import (
	"flag"
	"fmt"
)

// os.Args 获取命令行参数

func main() {
	// fmt.Println(os.Args)

	// name := flag.String("name", "xx", "请输入名字：")
	// age := flag.Int("age", 9000, "请输入真实年龄：")
	// married := flag.Bool("married", false, "结婚了吗")
	// cTime := flag.Duration("ct", time.Second, "结婚多久了")

	// 使用flag
	// flag.Parse()
	// fmt.Println(*name)
	// fmt.Println(*age)
	// fmt.Println(*married)
	// fmt.Println(*cTime)

	var name string
	flag.StringVar(&name, "name", "xx", "请输入名字：")
	flag.Parse()
	fmt.Println(name)
}
