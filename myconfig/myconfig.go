package main

import (
	"fmt"

	"myconfig/configs"
)

var cfg *configs.StudentConfig

func main() {
	cfg = &configs.StudentConfig{}
	configs.LoadJsonConfig(cfg, "settings.json")

	fmt.Println(cfg)
}
