package study_case

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// HttpPostCase 发送HTTP POST请求到指定URL，发送JSON数据并打印响应结果
// 该函数不接受任何参数
// 该函数没有返回值
func HttpPostCase() {
	// 定义请求URL、内容类型和要发送的JSON数据
	url := "https://www.runoob.com/try/ajax/demo_post2.php"
	contentType := "application/json"
	data := `{"fname": "RUNOOB","lname": "Boy"}`
	
	// 发送POST请求
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed ,err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	
	// 读取响应体内容
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed ,err:%v\n", err)
		return
	}

	// 打印响应结果
	fmt.Println(string(b))
}
