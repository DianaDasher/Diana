package test

import (
	"fmt"

	utils "github.com/DianaDasher/Diana/untils"
)

func main() {
	var appCfg AppConfig

	// 使用Cfg函数读取JSON配置文件
	err := utils.Cfg[AppConfig]("json", "./test/test.json", &appCfg)
	if err != nil {
		fmt.Printf("Failed to read settings: %s\n", err)
		return
	}

	global := appCfg
	fmt.Println(global)

	// 直接通过appCfg访问配置信息
	fmt.Println(">>>>>>>>>>>>")
	fmt.Printf("Server Port: %d\n", appCfg.Sev.Port)

	fmt.Printf("Database Host: %s\n", appCfg.DB.Host)
	fmt.Printf("Database port: %s\n", appCfg.DB.Port)
	fmt.Printf("Database name: %s\n", appCfg.DB.Name)
	fmt.Printf("Database user: %s\n", appCfg.DB.User)
	fmt.Printf("Database password: %s\n", appCfg.DB.Password)
	fmt.Println(">>>>>>>>>>>>")
	fmt.Printf("Server Port: %d\n", global.Sev.Port)

	fmt.Printf("Database Host: %s\n", global.DB.Host)
	fmt.Printf("Database port: %s\n", global.DB.Port)
	fmt.Printf("Database name: %s\n", global.DB.Name)
	fmt.Printf("Database user: %s\n", global.DB.User)
	fmt.Printf("Database password: %s\n", global.DB.Password)
}
