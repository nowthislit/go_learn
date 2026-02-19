// 练习2：文件复制与IO流
// 掌握 io 包和文件流操作

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("===== 练习2：文件复制与IO流 =====")

	// 准备测试文件
	srcPath := "source.txt"
	dstPath := "dest.txt"

	// TODO: 1. 创建源文件并写入一些内容
	os.Create(srcPath)
	os.WriteFile(srcPath, []byte("Welcome"), 0644)

	// TODO: 2. 使用 io.Copy 复制文件
	copyFile(srcPath, dstPath)

	// TODO: 3. 验证复制是否成功（比较内容）
	compareFiles(srcPath, dstPath)

	// TODO: 4. 实现一个复制大文件的函数，显示进度（每复制100字节打印一次）
	copyFileWithProgress(srcPath, dstPath)

	// TODO: 5. 清理测试文件
	os.Remove(srcPath)
}

// copyFile 使用 io.Copy 复制文件
func copyFile(src, dst string) (int64, error) {
	// TODO:
	// 1. 打开源文件
	// 2. 创建目标文件
	// 3. 使用 io.Copy 复制
	// 4. 关闭文件
	f, err := os.Open(src)
	tf, err := os.Create(dst)
	io.Copy(tf, f)
	defer f.Close()
	defer tf.Close()

	return 0, err
}

// copyFileWithProgress 复制文件并显示进度
func copyFileWithProgress(src, dst string) error {
	// TODO: 使用 io.Copy 的变体，或自己实现循环读取写入
	// 提示：使用 bufio.Reader 每次读取一定大小，打印进度
	f, _ := os.Open(src)
	d, _ := os.Create(dst)
	defer f.Close()
	defer d.Close()
	r := bufio.NewReader(f)
	max := 2
	info, _ := f.Stat()
	buff := make([]byte, max)
	for i := 0; i < int(info.Size())/max+1; i++ {
		n, _ := r.Read(buff)
		if n > 0 {
			d.Write(buff[:n])
			fmt.Printf("copying:%.f%%...\n", float64(max*i+n)/float64(int(info.Size()))*100)
		}
	}

	return nil
}

// compareFiles 比较两个文件内容是否相同
func compareFiles(file1, file2 string) (bool, error) {
	// TODO: 读取两个文件，比较内容
	cnt1, err := os.ReadFile(file1)
	cnt2, err := os.ReadFile(file2)
	if string(cnt1) == string(cnt2) {
		fmt.Println("文件内容一致")
		return true, nil
	}
	return false, err
}
