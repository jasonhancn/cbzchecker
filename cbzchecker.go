package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := getBaseDirectory()
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".cbz") {
			checkCbz(path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	fmt.Println("按回车退出")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func getBaseDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func checkCbz(path string) {
	zipFile, err := zip.OpenReader(path)
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()
	for _, file := range zipFile.File {
		if file.FileInfo().IsDir() {
			fmt.Println("格式错误：" + file.FileInfo().Name() + "是文件夹")
		} else if file.FileInfo().Size() == 0 {
			fmt.Println("发现空文件" + file.FileInfo().Name() + "，位于" + path)
		}
	}
}
