package study

import (
	"errors"
	"fmt"
)


// 创建错误 使用 errors.New 函数  err := errors.New("这是一个错误")



//函数通常在最后的返回值中返回错误信息，使用 errors.New 可返回一个错误信息
func Sqrt(f float64)(float64, error) {
	if f < 0 {
		return 0, errors.New("负数不能开方")
	}
	return f * f, nil //nil表示没有错误

}


// 显式返回错误



// 案例入口函数
func ErrorHandlingCase() {
	// 创建错误 使用 errors.New 函数
	err := errors.New("这是一个错误")
	fmt.Println(err)

	// 调用 Sqrt 函数
	result, err:= Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}


