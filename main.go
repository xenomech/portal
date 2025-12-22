package main

import (
	"fmt"
	"portal/config"
)

func main() {
	fmt.Println("hello world")
	cfg, err := config.Load()
	fmt.Println(cfg.Repos)
	fmt.Println(err)
	cfg.AddGitRepositoryToConfig(
		"abc",
		"abc",
	)
	cfg.Save()
}
