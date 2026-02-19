// 练习3：JSON基础
// 掌握 encoding/json 包的基础使用

package main

import (
	"encoding/json"
	"fmt"
	//"strings"
)

// User 用户结构体
type User struct {
	ID   int       `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
	Addr []Address `json:addrs`
}

type Address struct {
	Desc string `json:desc`
	No   int    `json:no`
}

func main() {
	fmt.Println("===== 练习3：JSON基础 =====")

	// TODO: 1. 创建一个 User 结构体实例
	user := User{
		ID:   1,
		Name: "joe",
		Age:  11,
	}

	// TODO: 2. 将结构体转换为 JSON 字符串（Marshal）
	jsonStr, _ := structToJSON(user)
	fmt.Println(jsonStr)

	// TODO: 3. 将 JSON 字符串转换回结构体（Unmarshal）
	user1 := User{}
	err := jsonToStruct(jsonStr, user1)
	if err == nil {
		fmt.Println(user)
	}

	// TODO: 4. 处理嵌套结构体（用户包含地址信息）
	user2 := User{
		ID:   1,
		Name: "joe",
		Age:  11,
		Addr: []Address{
			{
				Desc: "sss",
				No:   1,
			},
		},
	}

	jsonStr1, _ := structToJSON(user2)
	fmt.Println(jsonStr1)

	// TODO: 5. 处理数组（多个用户）
	users := [1]User{
		{
			ID:   1,
			Name: "joe",
			Age:  11,
			Addr: []Address{
				{
					Desc: "sss",
					No:   1,
				},
			},
		},
	}
	str, _ := structToJSON(users)
	// TODO: 6. 使用 MarshalIndent 美化输出
	printPrettyJSON(str)
}

// structToJSON 结构体转JSON
func structToJSON(v interface{}) (string, error) {
	// TODO: 使用 json.Marshal
	bs, err := json.Marshal(v)
	return string(bs), err
}

// jsonToStruct JSON转结构体
func jsonToStruct(jsonStr string, v interface{}) error {
	// TODO: 使用 json.Unmarshal
	return json.Unmarshal([]byte(jsonStr), &v)
}

// printPrettyJSON 美化打印JSON
func printPrettyJSON(v interface{}) {
	// TODO: 使用 json.MarshalIndent
	bs, _ := json.MarshalIndent(v, "", " ")
	fmt.Println(string(bs))
}
