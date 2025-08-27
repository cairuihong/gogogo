package study_case

import "fmt"

/*
组合接口
*/

type Reader interface {
	Read() string
}
type Writer interface {
	Write(data string)
}

type ReaderandWriter interface {
	Reader
	Writer
}

type File struct{}

func (f File) Read() string {
	return "Reading~~~~~~~"
}

func (f File) Write(data string) {
	fmt.Println("Writing data:", data)
}

func CoupleInterface() {
	var re ReaderandWriter = File{}
	fmt.Println(re.Read())
	re.Write("Hello~~~~~~~")
}
