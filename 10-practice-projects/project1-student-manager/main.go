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
	manager.AddStudent(Student{ID: 1, Name: "张三", Scores: []int{85, 90, 78}})

	// TODO: 2. 显示所有学生
	manager.ListStudents()

	// TODO: 3. 查询学生成绩
	student, err := manager.GetStudent(1)
	fmt.Println(student, err)

	// TODO: 4. 计算平均分
	avg := manager.CalculateAverage(1)
	fmt.Println(avg)
	// TODO: 5. 保存到文件
	manager.SaveToFile()

	// TODO: 6. 从文件加载
	manager.LoadFromFile()
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
	m.students = append(m.students, s)
}

// GetStudent 获取学生
func (m *StudentManager) GetStudent(id int) (*Student, error) {
	// TODO: 根据ID查找学生
	for _, stu := range m.students {
		if stu.ID == id {
			return &stu, nil
		}
	}
	return nil, fmt.Errorf("未找到")
}

// ListStudents 列出所有学生
func (m *StudentManager) ListStudents() {
	// TODO: 遍历并打印所有学生信息
	for _, stu := range m.students {
		fmt.Println(stu)
	}
}

// CalculateAverage 计算学生平均分
func (m *StudentManager) CalculateAverage(id int) float64 {
	// TODO: 找到学生，计算所有成绩的平均值
	for _, stu := range m.students {
		if stu.ID == id {
			var count float64
			for _, score := range stu.Scores {
				count += float64(score)
			}
			return count / float64(len(stu.Scores))
		}
	}
	return 0
}

// SaveToFile 保存到JSON文件
func (m *StudentManager) SaveToFile() error {
	// TODO: 使用 json.MarshalIndent 和 os.WriteFile
	bs, _ := json.MarshalIndent(m.students, "", " ")
	err := os.WriteFile(m.filename, bs, 0644)
	if err != nil {
		return err
	}
	return nil
}

// LoadFromFile 从JSON文件加载
func (m *StudentManager) LoadFromFile() error {
	// TODO: 使用 os.ReadFile 和 json.Unmarshal
	bs, _ := os.ReadFile(m.filename)
	err := json.Unmarshal(bs, &m.students)
	if err != nil {
		return err
	}
	fmt.Println(m.students)
	return nil
}
