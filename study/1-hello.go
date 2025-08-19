package main

import "fmt"

var ago int
var name string = "Alice"
var city = "New York"

func main1() {
	//school := "清华大学"
	fmt.Println("请输入你的姓名：")
	fmt.Scanln(&name)
	fmt.Println("你的名字是", name)
	// Print 基础打印，不换行
	fmt.Print(city)

	age := 01
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)
	fmt.Printf("你的年龄是 %d 岁 。\n", age)
}
