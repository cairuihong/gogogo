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

func FunctionsCase() {
	a := 100
	b := 200
	ret := max(a, b)
	fmt.Println(ret)

	c, d := swap("hello", "world")
	fmt.Println(c, d)

	println(calculator(a, b))

}
