package config

import (
	"github.com/spf13/viper"
)

var CfgParams configParameters

// configParameters 项目配置参数
type configParameters struct {
	WritevConfig writevConfig `mapstructure:"writevConfig"`
}

// writevConfig 集中写的配置项
type writevConfig struct {
	SrcDir   excelFile `mapstructure:"srcDir"`
	DestFile destFile  `mapstructure:"destFile"`
}

// excelFile excel文件
type excelFile struct {
	Path  string   `mapstructure:"path"`
	Sheet string   `mapstructure:"sheet"`
	Cells []string `mapstructure:"cells"`
}

// destFile 目标文件
type destFile struct {
	ExmPath  string `mapstructure:"exmPath"`
	DestPath string `mapstructure:"destPath"`
	Sheet    string `mapstructure:"sheet"`
	Cell     string `mapstructure:"cell"`
}

// Load 获取配置参数
func Load() error {
	//表示 先预加载匹配的环境变量
	viper.AutomaticEnv()

	// 从yaml文件获取nacos配置
	vconfig := viper.New()
	// 添加读取的配置文件路径
	vconfig.AddConfigPath("./")
	// 设置读取的配置文件
	vconfig.SetConfigName("writev")
	// 读取文件类型
	vconfig.SetConfigType("yaml")
	// 读取yaml
	err := vconfig.ReadInConfig()
	if err != nil {
		return err
	}
	// 转译yaml文件
	if err = vconfig.Unmarshal(&CfgParams); err != nil {
		return err
	}
	return nil
}
