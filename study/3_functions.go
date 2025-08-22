package study

import "fmt"

// 函数定义
func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func swap(x, y string) (string, string) {
	return y, x
}

// 命名了返回参数(mul int, div int) ，函数的return 可以裸返回
func calculator(a, b int) (mul int, div int) {
	mul = a * b
	if b != 0 {
		div = a / b
	}
	return // 裸返回
}

/*
闭包函数（匿名函数）
	定义：匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
匿名函数是一种没有函数名的函数，通常用于在函数内部定义函数，或者作为函数参数进行传递
	匿名函数的特点
	1、没有函数名
		使用 func 关键字直接定义函数体
		通过变量（如 add）来引用这个函数
	2、可以赋值给变量、
		可以像普通变量一样传递和使用
		通过变量名调用：add(1, 2)
	3、可以作为参数传递
*/

func getSequence() func() int {
	/*
		案例 getSequence 计数器
		getSequence()是一个返回函数的函数
		它声明了一个局部变量 i并初始化为 0
		返回一个匿名函数(在返回值写一个函数去操作外层函数的变量)，这个匿名函数：
		每次调用时将 i的值增加 1 (i+=1)
		返回增加后的 i值
		关键点：返回的匿名函数​​捕获​​了外部变量 i，形成了一个闭包
		【匿名函数是闭包的基础，但不是所有匿名函数都是闭包】
	*/
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func FunctionsCase() {
	a := 100
	b := 200
	ret := max(a, b)
	fmt.Println(ret)

	c, d := swap("hello", "world")
	fmt.Println(c, d)

	println(calculator(a, b))

	// 闭包函数实例调用
	// fmt.Fprintln(getSequence()) 错误写法，因为 getSequence() 返回的是一个函数，应该先调用这个函数再传值
	seq := getSequence()
	fmt.Println(seq()) //->1
	fmt.Println(seq()) //->2
	fmt.Println(seq()) //->3
	/*
		【为什么输出是1、2、3而不是1、1、1？？】
		闭包的工作原理
			1、变量捕获
				在 getSequence() 函数中，变量 i 只在函数首次调用时初始化为0
				返回的匿名函数捕获了这个变量 i，形成了闭包
				这个 i 变量在 getSequence() 函数执行完毕后仍然存在于内存中
			2、状态保持
				每次调用 seq() 时，操作的是同一个被捕获的变量 i
				第一次调用：i 从0变为1，返回1
				第二次调用：i 从1变为2，返回2
				第三次调用：i 从2变为3，返回3
			3、关键概念
				不会被重置为0，因为它不是在匿名函数内部定义的
				闭包使得匿名函数能够访问并修改外层函数的变量
				同一个 seq 变量持有对同一个 i 的引用
			这就是闭包的特性：函数与其词法环境的组合，使得函数可以访问其外部作用域中的变量，并且这些变量的状态会在多次调用之间保持。
	*/

	// 匿名函数的使用
	// 定义一个匿名函数 ，赋值给add
	add := func(a, b int) int {
		return a + b
	}
	// 调用匿名函数
	result := add(1, 2)
	fmt.Printf("计算1+2的值%d\n", result)

	//匿名函数作为参数传递
	// calculate 接收一个操作函数和两个整数参数，返回操作函数执行后的结果
	// operator: 接收两个整数参数并返回整数结果的函数
	// x: 参与运算的第一个整数
	// y: 参与运算的第二个整数
	// 返回值: 执行operator(x, y)后的结果
	calculate := func(operator func(int, int) int, x, y int) int {
		return operator(x, y)
	}
	
	// 使用calculate函数执行加法运算，计算3+4的值
	sum := calculate(add, 3, 4)
	fmt.Printf("计算3+4的值:%d\n", sum)


	// 也可以直接在函数调用中定义匿名函数
	difference := calculate(func(a, b int) int {
		return a - b
	}, 5, 2)
	fmt.Printf("计算5-2的值:%d\n", difference)
}
