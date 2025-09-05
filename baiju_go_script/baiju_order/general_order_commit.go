package baiju_order

import (
	"fmt"
	"gogogo/baiju_go_script/public_func"
)

var resUrl string = "https://iflow-sit.kaboss.cn/in-order-xxl/in/order/inOrderCommit"

func delResData(certificateName, phone, certificateNum, reportIp string) map[string]interface{} {
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
		"certificateName":    certificateName,
		"certificateNum":     certificateNum,
		"phone":              phone,
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
		"reportIp":           reportIp,
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
func GeneralOrderCommit() {
	fmt.Println("通用下单接口地址:", resUrl)
	certificateName := public_func.GetName(true)
	phone := public_func.GetPhone(true, true)
	certificateNum := public_func.GetCertificateNum(true)
	reportIp := public_func.GetRandomIPv4InChina()
	fmt.Println(delResData(certificateName, phone, certificateNum, reportIp))
	// public_func.PublicPost(resUrl, delResData(certificateName))
}
