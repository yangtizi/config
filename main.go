package main

import (
	"fmt"

	"github.com/yangtizi/config/jsonconfig"
	"github.com/yangtizi/config/yamlconfig"
)

// SConfig 配置
type SConfig struct {
	Redis string `yaml:"redis" json:"redis,omitempty"`
	MySQL string `yaml:"mysql" json:"mysql,omitempty"`

	Proxy struct {
		Host string `yaml:"host" json:"host"`
		Port int    `yaml:"port" json:"port"`
	} `yaml:"proxy" json:"proxy"`
}

var cfg *SConfig

func main() {
	loadConfigDemo()
}

func loadConfigDemo() {
	cfg = &SConfig{}
	jsonconfig.JSONParsing("./config.json", cfg)

	fmt.Println(cfg.Proxy.Host)

	fmt.Printf("加载的配置是cfg = [%v]\n", cfg)

	yamlconfig.YMALParsing("./config.yaml", cfg)
	fmt.Println(cfg.Proxy.Host)

	fmt.Printf("加载的配置是cfg2 = [%v]\n", cfg)
}
