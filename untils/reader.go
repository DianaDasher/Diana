package utils

import (
	"log"

	"github.com/spf13/viper"
)

/*
	读取配置文件的工具类
	参数：
			configType  	   	string : 读取的配置文件的类型，例如 "yml"
			configFilePath   	string : 读取的配置文件的位置，例如 "/resources/config.yml"
			target 	any	   : 要转化成的结构体，建议储传入结构体的指针
	返回：
			target 	any	   : 转化完的结构体，使用时要转化成对应的类。转化方式为： [target].(类型)
			error 		error   : 如果发生错误，返回错误信息
*/

func ReadSettings(configType string, configFilePath string, target interface{}) error {
	viper.AddConfigPath(".")
	viper.SetConfigType(configType)
	viper.SetConfigFile(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		log.Println("读取配置文件失败: %s\n", err.Error())
		return err
	}

	if err := viper.Unmarshal(&target); err != nil {
		log.Println("解析配置文件失败: %s\n", err.Error())
		return err
	} else {
		return nil
	}
}

/*
Cfg 初始化并返回配置对象，使用配置文件类型和路径来加载配置
configType 为配置文件的类型，如 "yaml"
参数 configPath 为配置文件的路径
参数 cfgPtr 为用于接收配置数据的结构体指针
返回值为配置结构体的指针和可能的错误
Cfg 初始化配置并返回提供的结构体指针，要求该结构体与配置文件内容匹配
*/
func Cfg[T any](configType string, configPath string, cfgPtr *T) error {
	// 读取配置
	var configValue T

	// 尝试读取配置文件并赋值给configValue
	err := ReadSettings(configType, configPath, &configValue)
	if err != nil {
		log.Println("初始化配置失败: %s\n", err.Error())
		return err
	}

	*cfgPtr = configValue
	return nil
}
