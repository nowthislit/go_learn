// 练习4：配置文件管理
// 使用 JSON 实现配置的读取、修改和保存

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config 应用配置
type Config struct {
	AppName  string   `json:"app_name"`
	Version  string   `json:"version"`
	Debug    bool     `json:"debug"`
	Database Database `json:"database"`
}

// Database 数据库配置
type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	fmt.Println("===== 练习4：配置文件管理 =====")

	configFile := "config.json"

	// TODO: 1. 创建默认配置并保存到文件
	cfg := createDefaultConfig()
	err := saveConfig(configFile, cfg)
	if err != nil {
		panic(err)
	}

	// TODO: 2. 从文件读取配置
	rCfg, err := loadConfig(configFile)
	if err != nil {
		panic(err)
	}

	// TODO: 3. 修改配置（如更改端口、开启Debug模式）
	pCfg := &rCfg
	pCfg.Debug = false
	pCfg.Database.Port = 3306

	// TODO: 4. 保存修改后的配置
	err = saveConfig(configFile, *pCfg)
	if err != nil {
		panic(err)
	}

	// TODO: 5. 验证保存成功
	config, err := loadConfig(configFile)
	if err != nil {
		panic(err)
	}
	printConfig(config)
	// TODO: 6. 清理配置文件
}

// createDefaultConfig 创建默认配置
func createDefaultConfig() Config {
	// TODO: 返回默认配置
	db := Database{
		Host:     "127.0.0.1",
		Port:     5432,
		Username: "postgres",
		Password: "root",
	}
	return Config{
		AppName:  "demo",
		Version:  "0.0.1",
		Debug:    true,
		Database: db,
	}
}

// saveConfig 保存配置到JSON文件
func saveConfig(filename string, config Config) error {
	// TODO: 使用 json.MarshalIndent 美化JSON，然后写入文件
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	bs, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(bs))
	if err != nil {
		return err
	}
	return nil
}

// loadConfig 从JSON文件加载配置
func loadConfig(filename string) (Config, error) {
	// TODO: 读取文件，解析JSON到Config
	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	cfg := &Config{}
	err = json.Unmarshal(data, cfg)
	if err != nil {
		return Config{}, err
	}
	return *cfg, nil
}

// printConfig 打印配置信息
func printConfig(config Config) {
	// TODO: 格式化打印配置内容
	fmt.Println(config)
}
