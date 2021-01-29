package main

import (
	"Flamingo/config"
	"Flamingo/server"
)

func main() {
	config.Load()

	// 装载路由
	r := server.NewRouter()
	_ = r.Run(":567")
}
