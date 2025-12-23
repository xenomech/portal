package main

import (
	"fmt"
	"portal/internal/config"
	"portal/internal/git"
)

func main() {
	fmt.Println("hello world")

	// Test GetStatus with portal repository
	status := git.GetStatus(`/Users/xenomech/_p/portal`)
	fmt.Printf("Status: %+v\n", status)

	cfg, err := config.Load()
	// fmt.Println(cfg.Repos)
	fmt.Println(err)
	cfg.AddGitRepositoryToConfig(
		"abc",
		"abc",
	)
	cfg.Save()
}
