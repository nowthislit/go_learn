// 练习3：学生成绩管理
// 使用 map 和结构体管理学生信息

package main

import "fmt"

func main() {
	fmt.Println("===== 练习3：学生成绩管理 =====")

	// 创建学生管理器
	manager := NewStudentManager()

	// 添加学生
	manager.AddStudent("张三", 90)
	manager.AddStudent("李四", 85)
	manager.AddStudent("王五", 78)

	// 查询学生
	score, err := manager.GetScore("张三")
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("张三的成绩: %d\n", score)
	}

	// 查询不存在的
	_, err = manager.GetScore("赵六")
	if err != nil {
		fmt.Println("查询赵六:", err)
	}

	// 计算平均分
	avg := manager.GetAverageScore()
	fmt.Printf("班级平均分: %.2f\n", avg)

	// 打印所有学生
	fmt.Println("\n所有学生:")
	manager.PrintAllStudents()
}

// StudentManager 学生成绩管理器
type StudentManager struct {
	scores map[string]int
}

// NewStudentManager 创建新的管理器
func NewStudentManager() *StudentManager {
	return &StudentManager{
		scores: make(map[string]int),
	}
}

// AddStudent 添加学生
func (sm *StudentManager) AddStudent(name string, score int) {
	sm.scores[name] = score
}

// GetScore 查询学生成绩
func (sm *StudentManager) GetScore(name string) (int, error) {
	score, ok := sm.scores[name]
	if !ok {
		return 0, fmt.Errorf("学生 %s 不存在", name)
	}
	return score, nil
}

// GetAverageScore 计算平均分
func (sm *StudentManager) GetAverageScore() float64 {
	if len(sm.scores) == 0 {
		return 0
	}

	total := 0
	for _, score := range sm.scores {
		total += score
	}

	return float64(total) / float64(len(sm.scores))
}

// PrintAllStudents 打印所有学生
func (sm *StudentManager) PrintAllStudents() {
	for name, score := range sm.scores {
		fmt.Printf("  %s: %d分\n", name, score)
	}
}
