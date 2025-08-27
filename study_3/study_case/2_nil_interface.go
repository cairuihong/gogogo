package study_case

import "fmt"

/*
空接口

空接口 interface{} 是 Go 的特殊接口，表示所有类型的超集。
无方法约束​​：空接口是未声明任何方法的接口，因此​​所有类型都自动满足其要求​​（隐式实现）。
​​动态类型存储​​：可存储任意数据类型（如 int、string、struct等），类似其他语言的 Object或 any类型
*/

// printValue 利用空接口 实现打印任意类型的值和类型信息
func printValue(val interface{}) {
	fmt.Printf("传入的值为：%v,类型为：%T\n", val, val)
}

/*
类型选择
type switch 是 Go 中的语法结构，用于根据接口变量的具体类型执行不同的逻辑。
*/

func printType(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Println("整型:", v)
	case string:
		fmt.Println("字符串:", v)
	case bool:
		fmt.Println("布尔值:", v)
	default:
		fmt.Println("未知类型:", v)
	}
}

func NilInterfaceCase() {
	printValue(42)
	printValue(11.22)
	printValue("hello")
	printValue([]int{2, 3})

	/*
	   利用空接口做类型断言
	   	类型断言用于从接口类型中提取其底层值
	   	为了避免 panic，可以使用带检查的类型断言：
	   	value, ok := iface.(Type)
	   		ok 是一个布尔值，表示断言是否成功。
	   		如果断言失败，value 为零值，ok 为 false。
	*/

	fmt.Println("类型断言案例：")
	//类型断言
	var i interface{} = 42
	_, ok := i.(string)
	if ok {
		fmt.Printf("断言成功，是string 类型")
	} else {
		fmt.Println("断言失败，不是string 类型")
	}

	//类型选择
	fmt.Println("类型选择案例：")
	printType(42)
	printType("hello")
	printType(true)
	printType(3.14)
}
