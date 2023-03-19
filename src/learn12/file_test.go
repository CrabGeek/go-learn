package learn12

import (
	"fmt"
	"os"
	"testing"
)

func TestFileInfo(t *testing.T) {
	path := "C:\\Users\\xgp98\\Desktop\\go\\go-learn\\src\\learn12\\hhh.txt"
	printMessage(path)
}

func printMessage(filePath string) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("err: ", err.Error())
	} else {
		fmt.Printf("数据类型是: %T \n", fileInfo)
		fmt.Println("文件名: ", fileInfo.Name())
		fmt.Println("是否为目录: ", fileInfo.IsDir())
		fmt.Println("文件大小: ", fileInfo.Size())
		fmt.Println("文件权限: ", fileInfo.Mode())
		fmt.Println("文件最后修改的时间: ", fileInfo.ModTime())
	}
}
