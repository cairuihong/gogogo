package study_case

/*
select 语句使得一个 goroutine 可以等待多个通信操作。
select 会阻塞，直到其中的某个 case 可以继续执行：

*/
import "fmt"

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func ChannelSelectCase() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}
