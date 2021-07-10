package main

import (
	"fmt"
	"reflect"
	"errors"
	"io/ioutil"
	"strings"
	"strconv"
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
	Host 	 string `ini:"host"`
	Port 	 int 	`ini:"port"`
	Password string `ini:"password"`
	Database int 	`ini:"database"`
}

type Config struct {
	MySQLConfig `ini:"mysql"`
	RedisConfig	`ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0. 参数的校验
	// 0.1 data参数必须是指针类型，因为需要在函数中对其赋值
	t := reflect.TypeOf(data)
	fmt.Println("xxx: ", t, t.Kind())
	if t.Kind() != reflect.Ptr{
		err = errors.New("data param should be a pointer")
		return err
	}
	// 0.2 data参数必须是结构体类型指针（配置文件中各种键值对需要赋值给结构体字段
	if t.Elem().Kind() != reflect.Struct{
		err = errors.New("data param should be a struct pointer")
		return err
	}
	// 1. 读文件得到字节类型的数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	lineSlice := strings.Split(string(b), "\n")
	// fmt.Println(lineSlice)
	// 2. 一行一行读取数据
	var structName string
	for index, line := range lineSlice {
		// 去掉字符串收尾的空格
		line = strings.TrimSpace(line)

		// 如果是空行，跳过
		if len(line) == 0 {
			continue
		}
		// 2.1 若是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2 若是是"["开头的就表示是节（section）
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", index + 1)
				return
			}
			// 去除首位的 [] ，得到内部的内容且去除首尾空格
			sectionName := strings.TrimSpace(line[1:len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", index + 1)
				return
			}
			// 根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找到对应的嵌套结构体，把字段名记下来
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}

		} else {
			// 2.3 如果不是[开头就是=分割的键值对
			// 1. 以等号分割这一行，等号左边是key，等号右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line: %d syntax error", index + 1)
				return
			}
			idx := strings.Index(line, "=")
			key := strings.TrimSpace(line[:idx])
			value := strings.TrimSpace(line[idx+1:])
			// 2. 根据structName 去data里面把对应的嵌套结构体给取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName)
			sType := sValue.Type()

			
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fieldType reflect.StructField
			// 3. 遍历嵌套结构体的没有给字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i)
				fieldType = field
				if field.Tag.Get("ini") == key {
					// 找到了对应的字段
					fieldName = field.Name
					break
				}
			}
			// 4. 如果key=tag，给这个字段赋值
			// 4.1 根据fieldName去取这个字段
			fileObj := sValue.FieldByName(fieldName)
			// 4.2 对其赋值
			fmt.Println(fieldName, fieldType.Type.Kind())
			switch fieldType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					return
				}
				fileObj.SetBool(valueBool)
			}
		}
	}
	
	return
}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil{
		fmt.Printf("load ini failed, err:%v", err)
		return
	}
	fmt.Printf("%#v\n", cfg)
}
