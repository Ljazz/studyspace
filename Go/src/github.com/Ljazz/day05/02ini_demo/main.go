package main

import (
	"fmt"
	"reflect"
	"errors"
	"ioutil"
)

// ini配置文件解析器

// MysqlConfig MySQL配置结构体
type MySQLConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host string `ini:"host"`
	Port int `ini:"port"`
	Password string `ini:"password"`
	Database int `ini:"database"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0. 参数的校验
	// 0.1 data参数必须是指针类型，因为需要在函数中对其赋值
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr{
		err = errors.New("data param should be a pointer")
		return
	}
	// 0.2 data参数必须是结构体类型指针（配置文件中各种键值对需要赋值给结构体字段
	if t.Elem().Kind() != reflect.Struct{
		err = errors.New("data param should be a struct pointer")
		return
	}
	// 1. 读文件得到字节类型的数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return 
	}
	lineSlice := strings.Split(string(b), "\n")
	fmt.Println(lineSlice)
	// 2. 一行一行读取数据
	// 2.1 若是注释就跳过
	// 2.2 若是是"["开头的就表示是节（section）
	// 2.3 如果不是[开头就是=分割的键值对
	
	return
}

func main() {
	var mc MySQLConfig
	err := loadIni("./conf.ini", &mc)
	if err != nil{
		fmt.Printf("load ini failed, err:%v", err)
	}
	fmt.Println(mc.Address, mc.Port, mc.Username, mc.Password)
}
