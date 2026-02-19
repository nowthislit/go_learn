// 项目1：学生成绩管理系统
// 结合知识点：结构体、方法、JSON、文件操作

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Student 学生
type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Scores []int  `json:"scores"`
}

// StudentManager 学生管理器
type StudentManager struct {
	students []Student
	filename string
}

func main() {
	fmt.Println("===== 项目1：学生成绩管理系统 =====")

	manager := NewStudentManager("students.json")

	// TODO: 1. 添加几个学生
	// manager.AddStudent(Student{ID: 1, Name: "张三", Scores: []int{85, 90, 78}})

	// TODO: 2. 显示所有学生
	// manager.ListStudents()

	// TODO: 3. 查询学生成绩
	// student, err := manager.GetStudent(1)

	// TODO: 4. 计算平均分
	// avg := manager.CalculateAverage(1)

	// TODO: 5. 保存到文件
	// manager.SaveToFile()

	// TODO: 6. 从文件加载
	// manager.LoadFromFile()
}

// NewStudentManager 创建管理器
func NewStudentManager(filename string) *StudentManager {
	return &StudentManager{
		students: []Student{},
		filename: filename,
	}
}

// AddStudent 添加学生
func (m *StudentManager) AddStudent(s Student) {
	// TODO: 添加到 students 切片
}

// GetStudent 获取学生
func (m *StudentManager) GetStudent(id int) (*Student, error) {
	// TODO: 根据ID查找学生
	return nil, fmt.Errorf("未找到")
}

// ListStudents 列出所有学生
func (m *StudentManager) ListStudents() {
	// TODO: 遍历并打印所有学生信息
}

// CalculateAverage 计算学生平均分
func (m *StudentManager) CalculateAverage(id int) float64 {
	// TODO: 找到学生，计算所有成绩的平均值
	return 0
}

// SaveToFile 保存到JSON文件
func (m *StudentManager) SaveToFile() error {
	// TODO: 使用 json.MarshalIndent 和 os.WriteFile
	return fmt.Errorf("未实现")
}

// LoadFromFile 从JSON文件加载
func (m *StudentManager) LoadFromFile() error {
	// TODO: 使用 os.ReadFile 和 json.Unmarshal
	return fmt.Errorf("未实现")
}
