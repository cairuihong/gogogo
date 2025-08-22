package study_case

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

func BasicDataStructuresCase() {

	//数组
	fmt.Println("数组练习")
	fmt.Println(numbers1)
	fmt.Println(numbers2)
	numbers3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(numbers3)
	fmt.Println(numbers3[0]) //1

	//切片
	fmt.Println("切片练习")
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

	// map 集合（字典）
	fmt.Println("***********map集合练习")
	// 用make 创建一个map ，此时map是空的
	map1 := make(map[string]int)
	fmt.Println(map1)
	// 向map中添加元素
	map1["a"] = 1
	map1["b"] = 2
	fmt.Println(map1) // 输出map[a:1 b:2]

	//也可以用map字面值的语法创建map ,直接指定初始值
	map2 := map[string]int{"c": 3, "d": 4}
	fmt.Println(map2)

	// 访问map： 通过key来获取value
	fmt.Println(map1["a"])
	fmt.Println(map2["c"])

	// 删除map中的元素 ,通过delete ，传map名称和key ，进行删除
	delete(map1, "a")
	fmt.Println(map1) //输出map[b:2]

	//x += y和x++等简短赋值语法也可以用在map上
	map1["b"]++
	fmt.Println(map1) // 输出map[b:3]

	map1["b"] = map1["b"] + 1
	fmt.Println(map1) // 输出map[b:4]

	// 遍历map  ，需要使用for 和range 进行遍历 ，遍历的结果是无序的，需要排序可以考虑使用sort
	fmt.Println("*******遍历map")
	map3 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(map3)
	for k, v := range map3 {
		fmt.Println(k, v)
	}
	//如果只需要遍历 key 或者value ，可以在for中使用下划线 _ ,表示忽略索引值
	for k, _ := range map3 {
		fmt.Println(k)
	}

	// 结构体 go 中的自定义数据集合，有点像面向对象编程的类，但不支持集成，更像组合
	// 结构体的定义  使用 type + struct
	fmt.Println("******结构体")
	type User struct {
		UserName    string
		Email       string
		SignInCount int
		IsActive    bool
	}
	// 结构体赋值 ,忽略的字段会被复制为0或空
	user1 := User{
		UserName:    "Alice",
		Email:       "alice@example.com",
		SignInCount: 1,
		IsActive:    true,
	}
	fmt.Println(user1)

	// 也可以按声明的字段顺序直接提供字段值赋值，但必须提供所有字段的值
	user2 := User{"Bob", "alice@example.com", 1, false}
	fmt.Println(user2)

	// 访问结构体字段 , 使用点号 . 操作符 [结构体.成员名]
	fmt.Println(user1.Email)
}
