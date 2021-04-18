package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// resp, err := http.Get("http://127.0.0.1:8000/xxx/?name=sb&age=18")
	// if err != nil {
	// 	fmt.Printf("get url failed, err:%v\n", err)
	// 	return
	// }
	data := url.Values{} // URL values
	urlObj, _ := url.Parse("http://127.0.0.1:8000/xxx/")
	data.Set("name", "马啸")
	data.Set("age", "20")
	queryStr := data.Encode() // URL encode之后的URL
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	if err != nil {
		fmt.Printf("create newRequest failed, err:%v\n", err)
		return
	}
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Printf("http.DefaultClient request failed, err:%v\n", err)
	// 	return
	// }
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("http.DefaultClient request failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	// 从resp中把服务端返回的数据读出来
	// var data []byte
	// resp.Body.Read()
	// resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
