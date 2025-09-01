package study_case

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func HttpGetCase() {

	// 1、使用 http.Get()快速发起简单请求
	/*
	   defer resp.Body.Close()确保资源释放
	   io.ReadAll 读取响应体
	*/

	// resp, err := http.Get("https://www.baidu.com")
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(body))

	//2、 带参数的 GET 请求
	//
	apiUrl := "https://www.runoob.com/"
	params := url.Values{}
	params.Add("s", "python 教程")

	//拼接URL参数
	fullURL := apiUrl + "?" + params.Encode()
	resp2, err := http.Get(fullURL) // 发起请求
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()

	body2, err := io.ReadAll(resp2.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body2))
}
