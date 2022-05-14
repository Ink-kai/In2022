package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	pwd := "E:\\Log"
	filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		fmt.Println(filepath.Join(path, info.Name())) //打印文件或目录名
		return nil
	})
}
