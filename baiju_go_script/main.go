package main

import "gogogo/baiju_go_script/baiju_order"

func main() {

	//参数说明 GeneralOrderCommit(需要下单的总订单数量,并发数，pageUrl<推广链接>)
	var pageUrl string = "https://wp.kaboss.cn/xxl/pre/pages/pick/index?id=1963529796221493250&channelId=1968257103217643522"

	baiju_order.GeneralOrderCommit(1, 1, pageUrl)

	// public_func.PublicDataCase()
}
