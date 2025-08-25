package study_case

import (
	"fmt"
	"log"
	"os"
)

/*
创建目录
1、创建单个目录 用os.Mkdir()
2、创建多级目录 用os.MkdirAll()
*/

func CreateDir() {
	// 创建单个目录
	err := os.Mkdir("newdir", 0755)
	if err != nil {
		fmt.Println(err)
	}

	// 创建多级目录
	err = os.MkdirAll("newdir/subdir", 0755)
	if err != nil {
		fmt.Println(err)
	}
}

/*
读取目录内容
os.ReadDir
*/
// ReadDir 读取当前目录内容并打印每个条目的信息
// 该函数会显示每个文件或目录的名称、大小和修改时间
func ReadDir() {
	// 读取当前目录下的所有条目
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err) // 如果读取失败则终止程序
	}

	// 遍历所有目录条目并打印详细信息
	for _, entry := range entries {
		info, _ := entry.Info() // 获取条目详细信息
		// 格式化输出: 名称(左对齐20字符) 大小(右对齐8字符) 修改时间(格式化为 YYYY-MM-DD HH:MM:SS)
		fmt.Printf("%-20s %8d %v\n",
			entry.Name(), info.Size(), info.ModTime().Format("2006-01-02 15:04:05"))
	}
}

/*
删除目录
1、删除空目录   os.Remove()
2、删除非空目录   os.RemoveAll()
*/
func DeleteDir() {
	//删除空目录
	err := os.Remove("emptydir")
	if err != nil {
		log.Fatal(err)
	}

	// 递归删除目录及其内容
	err2 := os.RemoveAll("path/to/dir")
	if err2 != nil {
		log.Fatal(err2)
	}
}


func DirCase() {
	// CreateDir()
	// ReadDir()
	// DeleteDir()
}
