package study_case

import (
	"bufio"
	"fmt"
	"os"
)

/*
文件创建
使用 os 包来创建文件
os.Create("文件名")
os.Create() 函数会创建一个文件并返回一个文件对象和一个错误信息
创建成功后，需要用 defer 关键字来确保文件在使用完毕后被关闭，即defer file.Close()
*/

func FileCreate() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Println("文件创建成功")
}

/*
文件的打开和关闭
打开： os.Open（）
	当需要指定打开模式和权限时，应使用os.OpenFile函数
关闭： file.Close()
*/

func FileOpenAndClose() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件打开成功")
	defer file.Close()
	fmt.Println("文件关闭成功")
}

/*
文件删除
os.Remove("文件名")
*/
func FileDelete() {
	err := os.Remove("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件删除成功")
}

/*
文件的读取
可以使用 bufio 包来逐行读取文件内容
bufio.NewScanner(file) 创建一个扫描器

可以用 os.ReadFile() 一次性读取整个文件内容
*/
func FileReadByBufio() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//创建一个扫描器，用for 逐行读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("文件读取完毕")
	defer file.Close()
}

func FileReadAllByOsReader() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	fmt.Println("整个文件读取完毕")
	defer file.Close()
}

/*文件的写入
可以用 bufio.NewWriter() 来创建一个带缓冲的写入器，逐行写入文件
也可以使用 os.WriteFile() 一次性写入整个文件内容
*/

func FileWriteByBufio() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(file) //创建一个带缓冲的写入器
	//逐行写入文件内容
	for i := 0; i < 5; i++ {
		_, err := writer.WriteString("这是第" + fmt.Sprint(i) + "行数据\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	writer.Flush() //将缓冲区的内容写入文件
	fmt.Println("文件写入完毕")
	defer file.Close()
}

func FileWriteByOsWriter() {
	data := []byte("这是使用 os.WriteFile 写入的内容\n")
	err := os.WriteFile("test.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("使用 os.WriteFile 写入文件成功")
}

/*
校验文件是否存在
os.Stat("文件名") 获取文件信息
os.IsNotExist(err) 判断文件是否存在
*/
func ChackFileExist() {
	if _, err := os.Stat("test.txt"); os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else {
		fmt.Println("文件存在")

	}
}

/*
文件重命名
使用 os.Rename("旧文件名", "新文件名")
*/
func ReNameFile() {
	err := os.Rename("test.txt", "new_test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件重命名成功")
	err2 := os.Rename("new_test.txt", "test.txt")
	if err != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("文件重命名成功")
}
func FileIoCase() {
	// FileCreate()
	// FileOpenAndClose()
	// FileDelete()
	// FileReadByBufio()
	// FileReadAllByOsReader()
	// FileWriteByBufio()
	// FileWriteByOsWriter()
	// ChackFileExist()
	ReNameFile()
}
