package study_case

import "fmt"

//6_pointer案例执行入口
func PointerCase() {
	fmt.Println("6_pointer案例执行入口")
	// Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
	var a int = 10
	fmt.Printf("变量的地址：%x\n", &a) // 输出变量 a 的地址->变量的地址：c000094070

	// 声明指针 var var_name *var-type  var [指针变量名]  *[指针类型]

	/*
		指针使用流程：
			1、定义指针变量。
			2、为指针变量赋值。
			3、访问指针变量中指向地址的值
	*/

	// 声明普通变量
	var b int = 10
	// 声明指针变量
	var ip *int
	// 为指针赋值
	ip = &b

	//打印变量b的内存地址
	fmt.Printf("变量b的内存地址为： %x\n", &b) //->c00000a0d0

	// 打印指针变量ip的内存地址
	fmt.Printf("指针变量ip的内存地址为： %x\n", ip) //->c00000a0d0

	// 打印指针变量ip所指地址的值
	fmt.Printf("变量ip所指地址的值为：%d\n", *ip) //->10

	/*
		被定义但没被赋值的指针，叫空指针 ，即 nil
	*/
	var p *int
	fmt.Printf("空指针p的值为：%v\n", p) //-><nil>

	/*
		指针数组 ： 一个数组里面每个元素都是指针
				定义方式： var array_name [size]*var-type
		数组指针 ： 一个指针指向某个数组
				定义方式： var ptr_name *[size]var-type
	*/
	var arr [3]int = [3]int{1, 2, 3}
	var arrptr [3]*int

	// 将数组arr的每个值赋值给指针数组 arrptr ，然后打印指针数组的内容
	for i := 0; i < len(arr); i++ {
		arrptr[i] = &arr[i]
		fmt.Printf("指针数组arrptr[%d] = %d\n", i, arrptr[i])
	}

}
