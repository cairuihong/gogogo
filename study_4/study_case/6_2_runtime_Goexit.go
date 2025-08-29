/*
退出当前协程
*/

package study_case

import (
	"fmt"
	"runtime"
)

func RuntimeGoexitCase() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")

			runtime.Goexit() // 退出当前协程，终止协程，跳过后续代码执行。  不会再打印 "C.defer"和"B"、"A"

			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {

	}
}
