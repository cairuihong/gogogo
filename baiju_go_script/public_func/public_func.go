package public_func

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func PublicPost(resUrl string, resData map[string]interface{}) error {
	jsonData, err := json.Marshal(resData)
	if err != nil {
		return fmt.Errorf("marshal failed: %v", err)
	}

	resp, err := http.Post(resUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("post failed: %v", err)
	}
	defer resp.Body.Close()

	// 打印接口入参
	fmt.Printf("接口请求参数: %s\n", string(jsonData))

	// 读取响应体内容
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("get resp failed: %v", err)
	}

	// 打印响应结果
	fmt.Printf("响应结果: %s", string(b))

	return nil
}

/*base64 编码*/
func Base64Encode(data string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(data)))
}
