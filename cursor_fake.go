package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// 在编译时通过 -ldflags 注入的版本信息
var (
	version = "dev"
)

// 获取存储路径
func getStoragePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		return ""
	}

	var basePath string
	switch runtime.GOOS {
	case "windows":
		basePath = filepath.Join(home, "AppData", "Roaming", "Cursor", "User", "globalStorage")
	case "darwin":
		basePath = filepath.Join(home, "Library", "Application Support", "Cursor", "User", "globalStorage")
	case "linux":
		basePath = filepath.Join(home, ".config", "Cursor", "User", "globalStorage")
	default:
		fmt.Println("Unsupported operating system")
		return ""
	}

	return filepath.Join(basePath, "storage.json")
}

// 生成随机的64位十六进制字符串
func generateRandomHex64() string {
	const chars = "0123456789abcdef"
	result := make([]byte, 64)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// 生成UUID式的字符串
func generateHexUUID() string {
	uuid := make([]byte, 32)
	for i := range uuid {
		uuid[i] = "0123456789abcdef"[rand.Intn(16)]
	}
	return fmt.Sprintf("%s-%s-4%s-%s-%s",
		string(uuid[0:8]),
		string(uuid[8:12]),
		string(uuid[13:16]),
		string(uuid[16:20]),
		string(uuid[20:32]))
}

// 显示版本信息
func showVersion() {
	fmt.Printf("Cursor Fake v%s\n", version)
	fmt.Println("A tool to reset Cursor IDE device IDs")
}

func main() {
	// 解析命令行参数
	versionFlag := flag.Bool("v", false, "显示版本信息")
	flag.Parse()

	// 如果指定了版本标志，显示版本信息并退出
	if *versionFlag {
		showVersion()
		return
	}

	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 获取存储路径
	storagePath := getStoragePath()
	if storagePath == "" {
		fmt.Println("Failed to get storage path")
		pressEnterToExit()
		return
	}

	// 读取文件
	data, err := os.ReadFile(storagePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		pressEnterToExit()
		return
	}

	// 解析JSON
	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		pressEnterToExit()
		return
	}

	// 生成新的ID
	newMacMachineId := generateRandomHex64()
	newMachineId := generateRandomHex64()
	newDevDeviceId := generateHexUUID()

	// 更新ID
	jsonData["telemetry.macMachineId"] = newMacMachineId
	jsonData["telemetry.machineId"] = newMachineId
	jsonData["telemetry.devDeviceId"] = newDevDeviceId

	// 转换为JSON并格式化
	newData, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
		pressEnterToExit()
		return
	}

	// 写回文件
	if err := os.WriteFile(storagePath, newData, 0644); err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		pressEnterToExit()
		return
	}

	// 打印结果
	fmt.Printf("\nCursor Fake v%s\n", version)
	fmt.Println("Success!")
	fmt.Printf("Path: %s\n", storagePath)
	fmt.Printf("New macMachineId: %s\n", newMacMachineId)
	fmt.Printf("New machineId: %s\n", newMachineId)
	fmt.Printf("New devDeviceId: %s\n", newDevDeviceId)

	pressEnterToExit()
}

func pressEnterToExit() {
	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln()
} 