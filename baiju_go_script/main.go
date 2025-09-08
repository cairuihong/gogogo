package main

import "gogogo/baiju_go_script/baiju_order"

func main() {

	//参数说明 GeneralOrderCommit(需要下单的总订单数量,并发数)
	baiju_order.GeneralOrderCommit(10, 5)

	// public_func.PublicDataCase()
}
