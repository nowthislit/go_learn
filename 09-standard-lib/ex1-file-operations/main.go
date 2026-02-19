// 练习1：文件基础操作
// 掌握 os 包的文件和目录操作

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("===== 练习1：文件基础操作 =====")

	dir := "test_dir"

	// TODO: 1. 创建目录 test_dir
	createDir(dir)
	f := dir + "/test.txt"
	// TODO: 2. 在目录中创建文件 test.txt，写入内容 "Hello, Go!"
	writeFile(f, "Hello, Go!")
	// TODO: 3. 读取文件内容并打印
	cnt, _ := readFile(f)
	fmt.Println(cnt)

	// TODO: 4. 获取文件信息（大小、修改时间等）
	getFileInfo(f)

	// TODO: 5. 列出目录下的所有文件
	listDir(dir)

	// TODO: 6. 重命名文件为 test_new.txt
	os.Rename(f, dir+"/test_new.txt")

	// TODO: 7. 删除文件和目录
	os.Remove(dir + "/test_new.txt")
	os.Remove(dir)

}

// createDir 创建目录
func createDir(path string) error {
	// TODO: 使用 os.Mkdir 或 os.MkdirAll
	return os.Mkdir(path, 0755)
}

// writeFile 写入文件
func writeFile(path string, content string) error {
	// TODO: 使用 os.WriteFile
	return os.WriteFile(path, []byte(content), 0644)
}

// readFile 读取文件
func readFile(path string) (string, error) {
	// TODO: 使用 os.ReadFile
	cnt, err := os.ReadFile(path)
	return string(cnt), err
}

// getFileInfo 获取文件信息
func getFileInfo(path string) error {
	// TODO: 使用 os.Stat 获取文件大小、修改时间等
	info, err := os.Stat(path)
	fmt.Println(info)
	return err
}

// listDir 列出目录内容
func listDir(path string) error {
	// TODO: 使用 os.ReadDir
	dirEntry, err := os.ReadDir(path)
	for entry := range dirEntry {
		fmt.Println(entry)
	}
	return err
}
