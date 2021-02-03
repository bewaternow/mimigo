package main

import (
	"mimigo/config"
	"mimigo/server"
)

func main() {
	config.Load()

	// 装载路由
	r := server.NewRouter()
	_ = r.Run(":567")
}
