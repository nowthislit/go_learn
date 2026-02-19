// 练习2：图书管理系统
// 设计 Book、Author、Library 等结构体

package main

import "fmt"

func main() {
	fmt.Println("===== 练习2：图书管理系统 =====")

	// TODO: 创建图书馆
	lib := NewLibrary("市图书馆")

	// TODO: 创建一些图书
	book1 := Book{Title: "乌合之众"}
	book2 := Book{Title: "我的奋斗"}

	// TODO: 添加图书到图书馆
	lib.AddBook(book1)
	lib.AddBook(book2)

	// TODO: 列出所有图书
	lib.ListBooks()

	// TODO: 查找特定图书
	book, found := lib.FindBook("乌合之众")
	if found {
		fmt.Printf("找到《%s》: %v\n", book.Title, found)
		// TODO: 借阅图书
		lib.BorrowBook("乌合之众")
	}
}

// Author 作者结构体
type Author struct {
	// TODO: 定义 Name 和 Email 字段
	Name  string
	Email string
}

// Book 图书结构体
type Book struct {
	// TODO: 定义 Title、ISBN、Author（Author类型）、IsBorrowed 字段
	Title      string
	ISBN       string
	Author     string
	IsBorrowed bool
}


// Library 图书馆结构体
type Library struct {
	// TODO: 定义 Name 和 Books 字段（Books 是 []Book 切片）
	Name  string
	Books []Book
}

// NewLibrary 创建新图书馆
func NewLibrary(name string) *Library {
	// TODO: 返回 Library 指针，初始化 Books 切片
	l := Library{Name: name, Books: []Book{}}
	return &l
}

// AddBook 添加图书到图书馆
func (l *Library) AddBook(book Book) {
	// TODO: 将 book 添加到 Books 切片
	l.Books = append(l.Books, book)
}

// ListBooks 列出所有图书
func (l *Library) ListBooks() {
	// TODO: 遍历 Books，打印每本书的信息
	fmt.Printf("图书馆: %s\n", l.Name)
	for _, book := range l.Books {
		fmt.Printf("书: %s\n", book.Title)
	}

}

// FindBook 根据书名查找图书
func (l *Library) FindBook(title string) (*Book, bool) {
	// TODO: 遍历 Books，查找匹配的 Title
	// 找到返回 (&book, true)，否则返回 (nil, false)
	for i, book := range l.Books {
		if book.Title == title {
			return &l.Books[i], true
		}
	}
	return nil, false
}

// BorrowBook 借阅图书
func (l *Library) BorrowBook(title string) bool {
	// TODO: 查找图书，如果存在且未借出，标记为已借出
	// 返回是否借阅成功
	for i, book := range l.Books {
		if book.Title == title && !book.IsBorrowed {
			l.Books[i].IsBorrowed = true
			return true
		}
	}
	return false
}

// ReturnBook 归还图书
func (l *Library) ReturnBook(title string) bool {
	// TODO: 查找图书，如果存在且已借出，标记为未借出
	for i, book := range l.Books {
		if book.Title == title && book.IsBorrowed {
			l.Books[i].IsBorrowed = false
			return false
		}
	}
	return true
}
