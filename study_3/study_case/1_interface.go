package study_case

import (
	"fmt"
	"math"
)

/*
接口的定义和实现
*/

// 定义一个接口 --形状（包含面积和周长方法）
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义一个结构体--圆形
type Circle struct {
	Radius float64
}

// Circle 实现Shape 接口
func (c Circle) Area() float64 {
	/*
		计算圆的面积
	*/
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	/*
		计算圆的周长
	*/
	return 2 * math.Pi * c.Radius
}

// 再定义一个结构体 -长方形
type Rectangle struct {
	Width  float64
	Height float64
}

// 计算 长方形的面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 计算长方形的周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func InterfaceCase() {
	c := Circle{Radius: 5}

	var s1 Shape = c
	// 接口变量 s 可以储存实现了 Shape 接口的任何值

	r := Rectangle{Width: 5, Height: 10}
	var s2 Shape = r

	fmt.Println("圆面积Area:", s1.Area())
	fmt.Println("圆周长Perimeter:", s1.Perimeter())
	fmt.Println("长方形面积Area:", s1.Area())
	fmt.Println("长方形周长Perimeter:", s2.Perimeter())

}
