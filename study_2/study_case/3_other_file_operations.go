package study_case

import (
	"io"
	"log"
	"os"
)

/*文件复制
io.Copy(目标文件, 源文件)
需要先打开文件
*/

func CopyFile() {
	srcFile, err := os.Open("srctest.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()
	// srcFile.Close()  defer只是注册了一个延迟调用，不会立即执行只会在函数返回前才被调用，所以才不会影响复制

	dstFile, err := os.Create("dsttest.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	bytesCopied, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("复制了 %d 字节\n", bytesCopied)

}

/*
文件内容追加
使用file.WriteString（）
*/
func FileAppend() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()
	if _, err := file.WriteString("追加新内容~~~！！！"); err != nil {
		log.Fatal(err)
	}
}

func OtherFileOperationsCase() {
	// CopyFile()
	FileAppend()
}
