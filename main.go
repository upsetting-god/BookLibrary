// main file
package main

import (
	"fmt"
	"main/core"
	"net"
	"os"

	"gopkg.in/yaml.v3"
)

func localip() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml: "server"`

	AllowedExtensions []string `yaml:"allowed_ex"`
}

var cfg *Config

func loadcfg() error {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return err
	}
	cfg = &Config{}

	return yaml.Unmarshal(data, cfg)
}

func main() {
	go core.Server()
	if err := loadcfg(); err != nil {
		fmt.Println("[CLIENT]: Error, can`t read config")
		return
	}
	banner := `
.__  ._____.
|  | |__\_ |______________ _______ ___.__.
|  | |  || __ \_  __ \__  \\_  __ <   |  |
|  |_|  || \_\ \  | \// __ \|  | \/\___  |
|____/__||___  /__|  (____  /__|   / ____|
             \/           \/       \/

`
	fmt.Println(banner)
	fmt.Println("Server was started")
	fmt.Printf("IP: %s	PORT: %d\n", localip(), cfg.Server.Port)

	select {}
}
