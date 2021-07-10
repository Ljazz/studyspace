package main

import (
	"fmt"
	"reflect"
)

type myInt int64

// json
type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}

func reflectType2(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) // 修改的是副本，reflect包会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem() 方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	// str := `{"name":"Tom", "age":9000}`
	// var p person
	// json.Unmarshal([]byte(str), &p)
	// fmt.Println(p.Name, p.Age)

	// var a float32 = 3.14
	// reflectType(a) // type:float32
	// var b int64 = 100
	// reflectType(b) // type:int64

	// var a *float32  // 指针
	// var b myInt     // 自定义类型
	// var c rune      // 类型别名
	// reflectType2(a) // type: kind:ptr
	// reflectType2(b) // type:myInt kind:int64
	// reflectType2(c) // type:int32 kind:int32

	// type student struct {
	// 	name string
	// 	age  int
	// }
	// type book struct{ title string }
	// var d = student{
	// 	name: "xiaowang",
	// 	age:  18,
	// }
	// var e = book{title: "Go"}
	// reflectType2(d) // type:student kind:struct
	// reflectType2(e) // type:book kind:struct

	// var a float32 = 3.14
	// var b int64 = 100
	// reflectValue(a) // type is float32, value is 3.140000
	// reflectValue(b) // type is int64, value is 100
	// c := reflect.ValueOf(10)
	// fmt.Printf("type c :%T\n", c) // type c :reflect.Value

	var a int64 = 100
	reflectSetValue1(a)
	// reflectSetValue2(&a)
	fmt.Println(a)
}
