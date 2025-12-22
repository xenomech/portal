package main

import (
	"fmt"
	"portal/config"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(config.ConfigDir())
	fmt.Println(config.ConfigPath())

}
