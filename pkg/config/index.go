package config

import (
	config "cutego/pkg/config/models"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var (
	// AppCoreConfig 核心配置
	AppCoreConfig *config.ApplicationCoreStruct
	// AppEnvConfig 环境配置
	AppEnvConfig *config.ApplicationEnvStruct
)

// GetRootPath 获取项目根路径
func GetRootPath() string {
	rootPath, _ := os.Getwd()
	return rootPath
}

// GetPathSeparator 获取路径分隔符
func GetPathSeparator() string {
	return string(os.PathSeparator)
}

// PathExists 判断文件或文件夹是否存在
// 如果返回的错误为nil,说明文件或文件夹存在
// 如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
// 如果返回的错误为其它类型,则不确定是否在存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// LoadYamlFile 加载yaml文件
func LoadYamlFile(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	err = yaml.Unmarshal(data, v)
	if err != nil {
		panic(err.Error())
	}
}

// BaseConfigDirPath 配置文件所在路径
const BaseConfigDirPath = "resources"

func readAppYmlFile(resourcePath string) {
	// 读取主配置文件
	applicationCoreFileName := BaseConfigDirPath + "/application.yml"
	applicationCoreFilePath := resourcePath + GetPathSeparator() + applicationCoreFileName
	exists, _ := PathExists(applicationCoreFilePath)
	if !exists {
		panic(applicationCoreFileName + "配置文件不存在!")
	}
	AppCoreConfig = &config.ApplicationCoreStruct{}
	// 由于要改变appConfig内部的值, 所以这里要取址
	LoadYamlFile(applicationCoreFilePath, AppCoreConfig)

	// 读取环境文件
	applicationEnvFileName := fmt.Sprintf(BaseConfigDirPath+"/application-%s.yml", AppCoreConfig.CuteGoConfig.Active)
	applicationEnvFilePath := resourcePath + GetPathSeparator() + applicationEnvFileName
	exists, _ = PathExists(applicationEnvFilePath)
	if !exists {
		panic(applicationEnvFileName + "配置文件不存在!")
	}
	AppEnvConfig = &config.ApplicationEnvStruct{}
	// 由于要改变appConfig内部的值, 所以这里要取址
	LoadYamlFile(applicationEnvFilePath, AppEnvConfig)
}

// ExecuteTip 执行日志
func ExecuteTip(message string) {
	fmt.Println(fmt.Sprintf("--------------------  %s  --------------------", message))
}

func init() {
	// 资源文件所在的路径
	resourcePath := GetRootPath()
	ExecuteTip("初始化应用配置 start")
	readAppYmlFile(resourcePath)
	ExecuteTip("初始化应用配置 end")
	return
}
