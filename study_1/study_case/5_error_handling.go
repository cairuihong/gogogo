package study_case

import (
	"errors"
	"fmt"
)

// 创建错误 使用 errors.New 函数  err := errors.New("这是一个错误")

// 函数通常在最后的返回值中返回错误信息，使用 errors.New 可返回一个错误信息
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("负数不能开方")
	}
	return f * f, nil //nil表示没有错误

}

// 显式返回错误   Go 中，错误通常作为函数的返回值返回
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

// 自定义错误
// 通过定义自定义类型，可以扩展 error 接口

/*
	 Go 语言中自定义错误类型的标准模式
	1、定义了一个自定义错误类型DivideError和一个除法函数divide2
	2、e 指向DivideError结构，实现了标准库的Error方法，
		表示DivideError类型可以直接作为 error使用
		当需要将 DivideError作为错误输出时，会调用此方法返回格式化的错误信息

	3、在divide2函数中使用DivideError来表示错误情况，
	通过 e.Dividend 和 e.Divisor 访问 DivideError 结构体的字段
*/

type DivideError struct {
	Dividend int
	Divisor  int
}

func (e *DivideError) Error() string {
	return fmt.Sprintf("无法将 %d 除以 %d", e.Dividend, e.Divisor)
}

func divide2(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivideError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}

// 案例入口函数
func ErrorHandlingCase() {
	// 创建错误 使用 errors.New 函数
	err := errors.New("这是一个错误")
	fmt.Println(err)

	// 调用 Sqrt 函数
	result, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// 显示返回错误
	result2, err2 := divide(10, 0)
	if err != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(result2)
	}

	// 自定义错误
	// 这里使用 _ 是因为 divide2 函数返回两个值，第一个是结果，第二个是错误,这里只需要错误，不需要结果
	_, err3 := divide2(10, 0)
	if err3 != nil {
		fmt.Println(err3)
	}

}
