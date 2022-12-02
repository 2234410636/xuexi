package main

import (
	"fmt"
	"os"
)

func main() {

	xxx := fileSize("C:/Users/yxy/Documents/WeChat Files/wxid_fho2nt452egp22/FileStorage/File/2022-11/校园网络使用说明.docx")
	fmt.Print(xxx)
}

func fileSize(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}
	// 延迟调用Close, 此时Close不会被调用
	//defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return 0
	}
	size := info.Size()
	f.Close()

	return size
	// defer机制触发, 调用Close关闭文件
}
