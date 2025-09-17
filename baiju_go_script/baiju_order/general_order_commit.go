package baiju_order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gogogo/baiju_go_script/public_func"
	"regexp"
	"sync"
	"time"
)

// 定义下单结果结构体
type OrderResult struct {
	Index   int         // 请求序号
	Success bool        // 是否成功
	Error   error       // 错误信息
	Data    interface{} // 响应数据
}

var resUrl string = "https://iflow-sit.kaboss.cn/in-order-xxl/in/order/inOrderCommit"

// 处理基础请求参数
func delResData() map[string]interface{} {
	var resData map[string]interface{} = map[string]interface{}{
		"sessionId":          public_func.GetUUID(),
		"iflowApi":           true,
		"iflowApiUrl":        true,
		"pageUrl":            "https://wp.kaboss.cn/xxl/sit/pages/pick/index?id=1963144332092182529&channelId=1963144383065559041",
		"pageId":             "1963144332092182529",
		"cacheId":            public_func.GetUUID(),
		"channelType":        1,
		"channelId":          "1963144383065559041",
		"buyPhone":           nil,
		"extData":            nil,
		"cacheFlag":          0,
		"certificateName":    nil,
		"certificateNum":     nil,
		"phone":              nil,
		"receivingAddress":   "5paw5riv5Lic5Y2X5Liw5rGHOTk55Y+3",
		"receivingProvince":  "广东省",
		"receivingCity":      "广州市",
		"receivingDistrict":  "海珠区",
		"frontPhoto":         nil,
		"frontPhotoSize":     nil,
		"behindPhoto":        nil,
		"behindPhotoSize":    nil,
		"headlessPhoto":      nil,
		"headlessPhotoSize":  nil,
		"dewuOrderId":        nil,
		"clickId":            "baiju",
		"ideaId":             "baiju",
		"mediaId":            "baiju",
		"planId":             "baiju",
		"projectId":          "baiju",
		"promotionId":        "baiju",
		"reportIp":           nil,
		"reportProvince":     "广东省",
		"reportCity":         "广州市",
		"feedback":           false,
		"packGroupId":        nil,
		"imagesMaterialId":   "",
		"titleMaterialId":    "",
		"videoMaterialId":    "",
		"advertiserId":       "",
		"csite":              "",
		"os":                 "",
		"txUrl":              "wp.kaboss.cn",
		"redirectUrl":        "",
		"code":               "",
		"mlId":               nil,
		"pageName":           "聚合-非选号-crh93",
		"isRedirect":         0,
		"orgPageId":          "",
		"orgChannelId":       "",
		"orgChannelType":     "",
		"channelRecordInfo":  []interface{}{},
		"orgOrderSourceType": "",
		"belongProvince":     nil,
		"belongCity":         nil,
		"agentId":            "",
		"openId":             "",
		"source":             "",
		"userAgent":          "mozilla/5.0 (windows nt 10.0; win64; x64) applewebkit/537.36 (khtml, like gecko) chrome/139.0.0.0 safari/537.36",
		"buttonName":         "",
		"commodityId":        "",
		"requestId":          public_func.GetUUID(),
		"businessCode":       "xxl",
	}
	return resData
}

// 生成订单参数函数
func generateOrderData(index int, pageUrl string) map[string]interface{} {
	// 基于原始参数创建新的订单参数
	orderData := make(map[string]interface{})
	for k, v := range delResData() {
		orderData[k] = v
	}

	// 定义部分参数(个性化参数)
	//进单渠道，产品信息
	orderData["pageUrl"] = pageUrl
	// 提取营销页ID
	idRegex := regexp.MustCompile(`[?&]id=([^&]*)`)
	idMatches := idRegex.FindStringSubmatch(pageUrl)
	if len(idMatches) > 1 {
		orderData["pageId"] = idMatches[1] // 营销页ID
	}
	// 提取渠道ID
	channelIdRegex := regexp.MustCompile(`[?&]channelId=([^&]*)`)
	channelIdMatches := channelIdRegex.FindStringSubmatch(pageUrl)
	if len(channelIdMatches) > 1 {
		orderData["channelId"] = channelIdMatches[1] // 渠道ID
	}

	// 用户信息
	orderData["certificateName"] = public_func.GetName(true)
	orderData["phone"] = public_func.GetPhone(true, true)
	orderData["certificateNum"] = public_func.GetCertificateNumByGenderAndAge(true, 1, 18)
	orderData["reportIp"] = public_func.GetRandomIPv4InChina()
	// 投放参数
	orderData["clickId"] = public_func.GetUUID()
	//地址参数
	address := public_func.GetAddress()
	orderData["receivingProvince"] = address.Province
	orderData["receivingCity"] = address.City
	orderData["receivingDistrict"] = address.District
	// fmt.Println(orderData)
	return orderData
}

// worker处理函数 (从通道中获取参数并执行下单操作)
func worker(requests <-chan map[string]interface{}, results chan<- OrderResult) {
	index := 0
	for req := range requests {
		data, err := public_func.PublicPost(resUrl, req)
		results <- OrderResult{
			Index:   index,
			Success: err == nil,
			Error:   err,
			Data:    data,
		}
		index++
	}
}

func GeneralOrderCommit(count int, concurrency int, pageUrl string) {
	/*
		count : 需要下单的总订单数量
		concurrency :并发数
	*/
	startTime := time.Now()
	fmt.Printf("计划下单数量：【%d】,并发数:【%d】\n", count, concurrency)
	fmt.Println("接口地址:", resUrl)

	// 创建带缓冲的channel用于控制并发数
	requests := make(chan map[string]interface{}, concurrency)
	results := make(chan OrderResult, count)

	var wg sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(requests, results)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	go func() {
		defer close(requests)
		for i := 0; i < count; i++ {
			orderData := generateOrderData(i, pageUrl) // 生成第i个订单参数
			requests <- orderData
		}
	}()

	successCount := 0
	failCount := 0

	for result := range results {
		if result.Success {
			successCount++
			jsonData, err := json.Marshal(result.Data)
			if err != nil {
				fmt.Println("JSON marshaling error:", err)
				continue
			}
			fmt.Printf(">>>下单成功,响应结果:%v\n", bytes.NewBuffer(jsonData))
		} else {
			failCount++
			fmt.Printf(">>>下单失败: %v\n", result.Error)
		}
	}
	//统计耗时
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	totalSeconds := int(duration.Seconds())
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60
	fmt.Printf(">>>>>>耗时：【%d小时%d分%d秒】\n", hours, minutes, seconds)

	fmt.Printf(">>>>>>批量下单任务完成 , 任务总数：【%d】 \n", successCount+failCount)
	fmt.Printf(">>>>>>结果：成功数量【%d】 , 失败数量【%d】\n", successCount, failCount)

}
