package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var serviceConfig ServiceConfig

func Config() *ServiceConfig {
	return &serviceConfig
}

type ServiceConfig struct {
	WorkDir      string       `json:"work_dir"`
	Version      string       `json:"version" yaml:"version"`
	Server       Server       `json:"server" yaml:"server"`
	Logger       LoggerConfig `json:"logger" yaml:"logger"`
	LoggerWriter LoggerWriter `json:"loggerwriter" yaml:"loggerwriter"`
	Redis        RedisConfig  `json:"redis" yaml:"redis"`
	Mysql        Mysql        `json:"mysql" yaml:"mysql"`
}

type Server struct {
	Name string `json:"name" yaml:"name"`
	Host string `json:"host" yaml:"host"`
	IP   string `json:"ip" yaml:"ip"`
	Port string `json:"port" yaml:"port"`
	Node int64  `json:"node" yaml:"node"`
}

func ParseConfig(path string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(path)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Read config file error: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		println(fmt.Sprintf("config file changed: %s", e.Name))
		PrettyPrint(Config())
		if err2 := v.Unmarshal(Config()); err2 != nil {
			fmt.Println(fmt.Errorf("Parse config file on changed error: %s \n", err2))
		}
	})
	if err1 := v.Unmarshal(Config()); err1 != nil {
		fmt.Println(fmt.Errorf("Parse config file error: %s \n", err1))
	}
	PrettyPrint(Config())

	return v
}

func PrettyPrint(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
		return
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		fmt.Println(v)
		return
	}

	fmt.Println(out.String())
}
