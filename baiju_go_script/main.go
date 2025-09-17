package main

import "gogogo/baiju_go_script/baiju_order"

func main() {

	//参数说明 GeneralOrderCommit(需要下单的总订单数量,并发数)  pageUrl(推广链接)
	var pageUrl string = "https://wp.kaboss.cn/xxl/sit/pages/pick/index?id=1950106606960062465&channelId=1950106665390911489"

	baiju_order.GeneralOrderCommit(300, 10, pageUrl)

	// public_func.PublicDataCase()
}
