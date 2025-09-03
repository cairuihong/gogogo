package baiju_order

import (
	"fmt"
	"gogogo/baiju_go_script/public_func"
)

var resUrl string = "https://iflow-sit.kaboss.cn/in-order-xxl/in/order/inOrderCommit"

// var resUrl string = "https://www.runoob.com/try/ajax/demo_post2.php"

// var resData map[string]interface{} = map[string]interface{}{
// 	"name":     "Go",
// 	"language": "C++",
// }

var resData map[string]interface{} = map[string]interface{}{
	"sessionId":          "df61e89a-2d70-438f-9bb5-67f350fe9991",
	"iflowApi":           true,
	"iflowApiUrl":        true,
	"pageUrl":            "https://wp.kaboss.cn/xxl/sit/pages/pick/index?id=1963144332092182529&channelId=1963144383065559041",
	"pageId":             "1963144332092182529",
	"cacheId":            "6877de2c-5069-4e3e-a2d3-026358b16772",
	"channelType":        1,
	"channelId":          "1963144383065559041",
	"buyPhone":           nil,
	"extData":            nil,
	"cacheFlag":          0,
	"certificateName":    "5rWL6K+V",
	"certificateNum":     "MTMwNzA1MjAxNzExMDUwMDcy",
	"phone":              "MTU4ODg4ODg4ODg=",
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
	"reportIp":           "223.160.225.152",
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
	"requestId":          "8a39c0c3-4c11-4e24-ac06-4cce81de2605",
	"businessCode":       "xxl",
}

func GeneralOrderCommit() {
	// url := "https://iflow-sit.kaboss.cn/in-order-xxl/in/order/inOrderCommit"

	fmt.Println("通用下单接口地址:", resUrl)
	// fmt.Println("请求参数:", resData)
	public_func.PublicPost(resUrl, resData)

}
