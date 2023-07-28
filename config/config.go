package config

import (
	"os"

	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	MySQL *MySQL `yaml:"mysql"`
}

type MySQL struct {
	UserName  string `yaml:"username"`
	Password  string `yaml:"password"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	DbName    string `yaml:"dbname"`
	Charset   string `yaml:"charset"`
	ParseTime string `yaml:"parse_time"`
	Loc       string `yaml:"loc"`
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err) //配置读取错误
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(err) //配置解析错误
	}
}
