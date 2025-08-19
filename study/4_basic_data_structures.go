package study

import "fmt"

//数组、切片、映射和结构体

// 数组声明需要指定元素类型及元素个数
// var arrayName [size]dataType
//var balance [10]float32  类型为float32 长度为10 的数组 balance

var numbers1 [5]int
var numbers2 = [5]int{1, 2, 3, 4, 5}

// 切片
// 声明一个未指定大小的数组来定义切片
// var identifier []type
// 定义切片
var s1 = []int{1, 2, 3}

// make函数创建切片
// 创建一个长度为3 ，容量为5 的切片，默认为[0,0,0]
var s3 = make([]int, 3, 5)

func case_4() {

	//数组
	fmt.Println(numbers1)
	fmt.Println(numbers2)
	numbers3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(numbers3)
	fmt.Println(numbers3[0]) //1

	//切片
	s2 := []int{4, 5, 6}
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(s3)
	//修改切片的值
	s3[0] = 100
	fmt.Println(s3)
	//追加切片的值
	s3 = append(s3, 400, 500, 600, 700, 800, 900, 1000, 1100)
	fmt.Println(s3)
	fmt.Println(cap(s3))

	//nil 切片 空切片
	var s4 []int
	fmt.Println(s4)
	if s4 == nil {
		fmt.Println("nil!!")
	} else {
		fmt.Println("not nil!!")
	}

	// 切片copy
	s5 := []int{1, 2, 3, 4, 5}
	var s6 = s5
	var s7 []int
	copy(s7, s5)
	fmt.Println(s5, s6, s7)
}
