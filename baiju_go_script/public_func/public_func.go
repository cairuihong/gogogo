package public_func

import (
	"bytes"
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

	// 读取响应体内容
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("get resp failed: %v", err)
	}

	// 打印响应结果
	fmt.Printf("响应结果: %s", string(b))

	return nil
}
